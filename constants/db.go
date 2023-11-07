package constants

import "time"

const DB = "portal"
const CONNECTION_STRING = "ATLAS_URI"

const COLLECTION_EXAM = "exams"
const COLLECTION_STUDENT = "students"
const COLLECTION_ROOM = "rooms"
const COLLECTION_ATTENDANCE = "attendances"

const FIREBASE_PROJECT_ID = "FIREBASE_PROJECT_ID"

const REDIS_HOST = "REDIS_HOST"
const REDIS_CACHING_LIMIT = 100000
const REDIS_CACHING_DURATION = time.Hour * 24 * 30

// Keys of Redis
const REDIS_GCP_JWKS = "GCP_JWKS"
