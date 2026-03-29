// types.ts
import { v4 as uuidv4 } from 'uuid';

export type FileMetadata = {
  uuid: string;
  name: string;
  size: number;
  mimeType: string;
  lastModified: Date;
};

export type ParsedData = {
  id: string;
  file: FileMetadata;
  parsedFields: string[];
  errors: string[];
};

export type DataParserOptions = {
  maxConcurrentRequests: number;
  retryTimeout: number;
  maxRetryCount: number;
};

export type File = {
  id: string;
  name: string;
  size: number;
  mimeType: string;
  lastModified: Date;
  binaryContent: Buffer;
};

export type ParsedFile = {
  id: string;
  file: FileMetadata;
  parsedFields: string[];
  errors: string[];
  binaryContent: Buffer;
};