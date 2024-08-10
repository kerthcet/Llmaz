/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package objectStore

import (
	"context"
	"io"
	"net/http"
	"os"
	"strconv"

	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

// ossConfig stores the configuration for oss bucket.
type ossConfig struct {
	Endpoint        string `yaml:"endpoint"`
	Bucket          string `yaml:"bucket"`
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}

// OSSBucket implements the store.Bucket interface.
type OSSBucket struct {
	name   string
	client *alioss.Client
	bucket *alioss.Bucket
}

func NewOSSBucket(ctx context.Context) (*OSSBucket, error) {
	config := ossConfig{
		Endpoint:        os.Getenv("ALIYUNOSS_ENDPOINT"),
		Bucket:          os.Getenv("ALIYUNOSS_BUCKET"),
		AccessKeyID:     os.Getenv("ALIYUNOSS_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("ALIYUNOSS_ACCESS_KEY_SECRET"),
	}

	if err := validate(config); err != nil {
		return nil, err
	}

	client, err := alioss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		return nil, errors.Wrap(err, "create aliyun oss client failed")
	}
	bk, err := client.Bucket(config.Bucket)
	if err != nil {
		return nil, errors.Wrapf(err, "use aliyun oss bucket %s failed", config.Bucket)
	}

	bkt := &OSSBucket{
		name:   config.Bucket,
		client: client,
		bucket: bk,
	}
	return bkt, nil
}

func (b *OSSBucket) Name() string {
	return b.name
}

// Get returns a reader for the given object name.
func (b *OSSBucket) Get(ctx context.Context, name string) (io.ReadCloser, error) {
	return b.getRange(ctx, name, 0, -1)
}

// Upload the contents of the reader as an object into the bucket.
func (b *OSSBucket) Upload(_ context.Context, name string, r io.Reader) error {
	panic("not implemented")
}

// Delete removes the object with the given name.
func (b *OSSBucket) Delete(ctx context.Context, name string) error {
	panic("not implemented")
}

func (b *OSSBucket) Close() error { return nil }

func (b *OSSBucket) getRange(_ context.Context, name string, off, length int64) (io.ReadCloser, error) {
	if name == "" {
		return nil, errors.New("given object name should not empty")
	}

	var opts []alioss.Option
	if length != -1 {
		opt, err := b.setRange(off, off+length-1, name)
		if err != nil {
			return nil, err
		}
		opts = append(opts, opt)
	}

	resp, err := b.bucket.GetObject(name, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *OSSBucket) setRange(start, end int64, name string) (alioss.Option, error) {
	var opt alioss.Option
	if 0 <= start && start <= end {
		header, err := b.bucket.GetObjectMeta(name)
		if err != nil {
			return nil, err
		}

		size, err := strconv.ParseInt(header["Content-Length"][0], 10, 64)
		if err != nil {
			return nil, err
		}

		if end > size {
			end = size - 1
		}

		opt = alioss.Range(start, end)
	} else {
		return nil, errors.Errorf("Invalid range specified: start=%d end=%d", start, end)
	}
	return opt, nil
}

// exists checks if the given object exists in the bucket.
func (b *OSSBucket) exists(ctx context.Context, name string) (bool, error) {
	exists, err := b.bucket.IsObjectExist(name)
	if err != nil {
		if b.isObjNotFoundErr(err) {
			return false, nil
		}
		return false, errors.Wrap(err, "cloud not check if object exists")
	}

	return exists, nil
}

// IsObjNotFoundErr returns true if error means that object is not found. Relevant to Get operations.
func (b *OSSBucket) isObjNotFoundErr(err error) bool {
	switch aliErr := errors.Cause(err).(type) {
	case alioss.ServiceError:
		if aliErr.StatusCode == http.StatusNotFound {
			return true
		}
	}
	return false
}

// validate checks to see the config options are set.
func validate(config OSSConfig) error {
	if config.Endpoint == "" || config.Bucket == "" {
		return errors.New("aliyun oss endpoint or bucket is not present in config file")
	}
	if config.AccessKeyID == "" || config.AccessKeySecret == "" {
		return errors.New("aliyun oss access_key_id or access_key_secret is not present in config file")
	}

	return nil
}
