/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package emailidentity

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/sesv2"
	svcsdk "github.com/aws/aws-sdk-go/service/sesv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/sesv2/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an EmailIdentity resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create EmailIdentity in AWS"
	errUpdate        = "cannot update EmailIdentity in AWS"
	errDescribe      = "failed to describe EmailIdentity"
	errDelete        = "failed to delete EmailIdentity"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.EmailIdentity)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.EmailIdentity)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetEmailIdentityInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetEmailIdentityWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateEmailIdentity(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.EmailIdentity)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateEmailIdentityInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateEmailIdentityWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.DkimAttributes != nil {
		f0 := &svcapitypes.DkimAttributes{}
		if resp.DkimAttributes.CurrentSigningKeyLength != nil {
			f0.CurrentSigningKeyLength = resp.DkimAttributes.CurrentSigningKeyLength
		}
		if resp.DkimAttributes.LastKeyGenerationTimestamp != nil {
			f0.LastKeyGenerationTimestamp = &metav1.Time{*resp.DkimAttributes.LastKeyGenerationTimestamp}
		}
		if resp.DkimAttributes.NextSigningKeyLength != nil {
			f0.NextSigningKeyLength = resp.DkimAttributes.NextSigningKeyLength
		}
		if resp.DkimAttributes.SigningAttributesOrigin != nil {
			f0.SigningAttributesOrigin = resp.DkimAttributes.SigningAttributesOrigin
		}
		if resp.DkimAttributes.SigningEnabled != nil {
			f0.SigningEnabled = resp.DkimAttributes.SigningEnabled
		}
		if resp.DkimAttributes.Status != nil {
			f0.Status = resp.DkimAttributes.Status
		}
		if resp.DkimAttributes.Tokens != nil {
			f0f6 := []*string{}
			for _, f0f6iter := range resp.DkimAttributes.Tokens {
				var f0f6elem string
				f0f6elem = *f0f6iter
				f0f6 = append(f0f6, &f0f6elem)
			}
			f0.Tokens = f0f6
		}
		cr.Status.AtProvider.DkimAttributes = f0
	} else {
		cr.Status.AtProvider.DkimAttributes = nil
	}
	if resp.IdentityType != nil {
		cr.Status.AtProvider.IdentityType = resp.IdentityType
	} else {
		cr.Status.AtProvider.IdentityType = nil
	}
	if resp.VerifiedForSendingStatus != nil {
		cr.Status.AtProvider.VerifiedForSendingStatus = resp.VerifiedForSendingStatus
	} else {
		cr.Status.AtProvider.VerifiedForSendingStatus = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.EmailIdentity)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteEmailIdentityInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteEmailIdentityWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.SESV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.SESV2API
	preObserve     func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.GetEmailIdentityInput) error
	postObserve    func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.GetEmailIdentityOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.EmailIdentityParameters, *svcsdk.GetEmailIdentityOutput) error
	isUpToDate     func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.GetEmailIdentityOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.CreateEmailIdentityInput) error
	postCreate     func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.CreateEmailIdentityOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.DeleteEmailIdentityInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.EmailIdentity, *svcsdk.DeleteEmailIdentityOutput, error) error
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.EmailIdentity, *svcsdk.GetEmailIdentityInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.EmailIdentity, _ *svcsdk.GetEmailIdentityOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.EmailIdentityParameters, *svcsdk.GetEmailIdentityOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.EmailIdentity, *svcsdk.GetEmailIdentityOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.EmailIdentity, *svcsdk.CreateEmailIdentityInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.EmailIdentity, _ *svcsdk.CreateEmailIdentityOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.EmailIdentity, *svcsdk.DeleteEmailIdentityInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.EmailIdentity, _ *svcsdk.DeleteEmailIdentityOutput, err error) error {
	return err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
