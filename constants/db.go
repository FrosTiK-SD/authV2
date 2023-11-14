package constants

import "time"

const DB = "portal"
const CONNECTION_STRING = "ATLAS_URI"

const COLLECTION_EXAM = "exams"
const COLLECTION_STUDENT = "students"
const COLLECTION_GROUP = "groups"
const COLLECTION_ROOM = "rooms"
const COLLECTION_ATTENDANCE = "attendances"

const FIREBASE_PROJECT_ID = "FIREBASE_PROJECT_ID"

const CACHING_DURATION = 20 * time.Minute
const CACHE_CONTROL_HEADER = "Cache-Control"
const NO_CACHE = "no-cache"

// Keys of Cache
const GCP_JWKS = "GCP_JWKS"

// Operation Types
var DB_FINDONE = "DB_FINDONE"
var DB_FIND = "DB_FIND"
