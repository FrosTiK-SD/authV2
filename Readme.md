1. Go stores empty arrays as null by default when you assign an empty array and store it in mongodb
2. Omitempty doesnot work on structs
3. Our datetimes are in Unix Time Stamp in UTC Timezone
4. Synchronise your computer clock if your iat in JWT is causing issues.