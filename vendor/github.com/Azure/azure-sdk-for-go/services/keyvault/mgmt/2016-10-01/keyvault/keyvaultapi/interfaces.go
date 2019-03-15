package keyvaultapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

// VaultsClientAPI contains the set of methods on the VaultsClient type.
type VaultsClientAPI interface {
	CheckNameAvailability(ctx context.Context, vaultName keyvault.VaultCheckNameAvailabilityParameters) (result keyvault.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters keyvault.VaultCreateOrUpdateParameters) (result keyvault.Vault, err error)
	Delete(ctx context.Context, resourceGroupName string, vaultName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.Vault, err error)
	GetDeleted(ctx context.Context, vaultName string, location string) (result keyvault.DeletedVault, err error)
	List(ctx context.Context, top *int32) (result keyvault.ResourceListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.VaultListResultPage, err error)
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
	ListDeleted(ctx context.Context) (result keyvault.DeletedVaultListResultPage, err error)
	PurgeDeleted(ctx context.Context, vaultName string, location string) (result keyvault.VaultsPurgeDeletedFuture, err error)
	Update(ctx context.Context, resourceGroupName string, vaultName string, parameters keyvault.VaultPatchParameters) (result keyvault.Vault, err error)
	UpdateAccessPolicy(ctx context.Context, resourceGroupName string, vaultName string, operationKind keyvault.AccessPolicyUpdateKind, parameters keyvault.VaultAccessPolicyParameters) (result keyvault.VaultAccessPolicyParameters, err error)
}

var _ VaultsClientAPI = (*keyvault.VaultsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result keyvault.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*keyvault.OperationsClient)(nil)
