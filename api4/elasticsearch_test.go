// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"testing"
)

func TestElasticsearchTest(t *testing.T) {
	th := Setup().InitBasic().InitSystemAdmin()
	defer TearDown()

	_, resp := th.Client.TestElasticsearch()
	CheckForbiddenStatus(t, resp)

	_, resp = th.SystemAdminClient.TestElasticsearch()
	CheckNotImplementedStatus(t, resp)
}
