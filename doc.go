/*
Package isprod is a simple package to check if the application is running in production or not.

It has default conditions that fit the most common cases, but you can also add your own conditions.

Default rules are:

	If environment variable 'prod' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'production' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'staging' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'live' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'ci' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'PROD' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'PRODUCTION' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'STAGING' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'LIVE' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'CI' is set and its value is not one of [false], consider it as production environment.
	If environment variable 'env' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
	If environment variable 'environment' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
	If environment variable 'mode' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
	If environment variable 'ENV' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
	If environment variable 'ENVIRONMENT' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
	If environment variable 'MODE' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
*/
package isprod
