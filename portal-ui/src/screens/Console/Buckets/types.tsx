// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

export interface BucketEncryptionInfo {
  algorithm: string;
  kmsMasterKeyID: string;
}

export interface Details {
  tags?: object;
  locking?: boolean;
  quota?: object;
  versioning?: boolean;
}

export interface BucketInfo {
  name: string;
  access: string;
  definition: string;
  creation_date?: string;
  objects?: number;
  size?: number;
}

export interface BucketEvent {
  id: string;
  arn: string;
  events: string[];
  prefix: string;
  suffix: string;
}

export interface BucketEventList {
  events: BucketEvent[];
  total: number;
}

export interface ArnList {
  arns: string[];
}

export interface BucketVersioningInfo {
  excludeFolders?: boolean;
  excludedPrefixes?: Record<"Prefix", string>[];
  MFADelete?: string;
  status?: "Enabled" | "Suspended" | "";
}

export interface BucketObjectLocking {
  object_locking_enabled: boolean;
}

export interface BucketReplicationDestination {
  bucket: string;
}

export interface BucketReplicationRule {
  id: string;
  status: string;
  priority: number;
  delete_marker_replication: boolean;
  deletes_replication: boolean;
  metadata_replication: boolean;
  prefix?: string;
  tags?: string;
  destination: BucketReplicationDestination;
  syncMode: string;
  storageClass?: string;
  existingObjects?: boolean;
}

export interface BucketReplication {
  rules: BucketReplicationRule[];
}

export interface BucketQuota {
  quota: number;
  type: string;
}

export interface ChangePasswordRequest {
  current_secret_key: string;
  new_secret_key: string;
}

export interface ChangeUserPasswordRequest {
  selectedUser: string;
  newSecretKey: string;
}

export interface BulkReplicationResponse {
  replicationState: BulkReplicationItem[];
}

export interface BulkReplicationItem {
  errorString: string;
  originBucket: string;
  targetBucket: string;
}

interface IExpirationLifecycle {
  days: number;
  date: string;
  delete_marker?: boolean;
  noncurrent_expiration_days?: number;
}

interface ITransitionLifecycle {
  days: number;
  date: string;
  storage_class?: string;
  noncurrent_transition_days?: number;
  noncurrent_storage_class?: string;
}

export interface LifeCycleItem {
  id: string;
  prefix?: string;
  expiration?: IExpirationLifecycle;
  transition?: ITransitionLifecycle;
  tags?: any;
  status?: string;
}

export interface MultiBucketResult {
  bucketName: string;
  error?: string;
}

export interface MultiBucketResult {
  results: MultiBucketResult[];
}

export interface ITiersDropDown {
  label: string;
  value: string;
}
