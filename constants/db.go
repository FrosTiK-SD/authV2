package constants

import "time"

const DB = "portal"
const CONNECTION_STRING = "ATLAS_URI"
const REDIS_URI = "REDIS_URI"
const REDIS_USERNAME = "REDIS_USERNAME"
const REDIS_PASSWORD = "REDIS_PASSWORD"

const COLLECTION_EXAM = "exams"
const COLLECTION_STUDENT = "students"
const COLLECTION_GROUP = "groups"
const COLLECTION_ROOM = "rooms"
const COLLECTION_ATTENDANCE = "attendances"
const COLLECTION_RECRUITER = "recruiters"
const COLLECTION_DOMAIN = "domains"
const COLLECTION_COMPANY = "companies"

const FIREBASE_PROJECT_ID = "FIREBASE_PROJECT_ID"

const CACHING_DURATION = 20 * time.Hour
const CACHE_CONTROL_HEADER = "cache-control"
const NO_CACHE = "no-cache"

// Keys of Cache
const GCP_JWKS = "GCP_JWKS"
