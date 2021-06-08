<?php
$settings['db_host'] = $_ENV['CHEVERETO_DB_HOST'];
$settings['db_port'] = $_ENV['CHEVERETO_DB_PORT'];
$settings['db_name'] = $_ENV['CHEVERETO_DB_NAME'];
$settings['db_user'] = $_ENV['CHEVERETO_DB_USERNAME'];
$settings['db_pass'] = $_ENV['CHEVERETO_DB_PASSWORD'];
$settings['db_table_prefix'] = $_ENV['CHEVERETO_DB_PREFIX'];
$settings['db_driver'] = 'mysql';
$settings['db_pdo_attrs'] = [];
$settings['debug_level'] = 1;

// Advanced configurations
$val = getenv('CHEVERETO_SESSION_SAVE_PATH');
if ($val !== false ) {
    /*
     * Specify where Chevereto can save the user session. Combining this with a
     * shared volume/EFS/File synching, you can set up a highly available
     * cluster of Free Chevereto while ensure user login status are not lost if
     * they reconnect to another container
     */
    $settings['session.save_path'] = $val;
}

$val = getenv('CHEVERETO_ERROR_REPORTING');
if ($val !== false) {
    /*
     * Specify whether we wish to have error reporting or not
     * Note that while environment variable can only be string, Chevereto expect
     * this to be boolean - https://github.com/Chevereto/Chevereto-Free/blob/0184f27a97daa55ec3b07560c5dd619d22abc907/lib/G/G.php#L95
     * so here we have to convert the string true/false to the boolean value
     */
    $settings['error_reporting'] = strtolower(trim($val)) === 'true';
}

$val = getenv('CHEVERETO_DEFAULT_TIMEZONE');
if ($val !== false) {
    /*
     * Specify the  default timezone for any newly created/registered user
     * Please make sure to use a PHP supported timezone - https://www.php.net/manual/en/timezones.php - eg America/New_York
     */
    $settings['default_timezone'] = $val;
}

$val = getenv('CHEVERETO_ENVIRONMENT');
if ($val !== false) {
    /*
     * Specify the environment of G\ app
     * This is usually useful for debugging purpose - see https://github.com/Chevereto/G-Library/blob/da673a3abef80ad9ac1aa3f26fb7abe4973c8445/lib/G/functions.php#L1649
     */
    $settings['error_reporting'] = $val;
}

$val = getenv('CHEVERETO_HTTPS');
if ($val !== false) {
    /*
     * Specify whether user should connect to Chevereto via HTTPS instead of HTTP
     * Please note that Chevereto has excellent mechanism to detect this
     * already - https://github.com/Chevereto/Chevereto-Free/blob/0184f27a97daa55ec3b07560c5dd619d22abc907/lib/G/G.php#L113
     * So 99% of the time you won't have to bother with this. The only
     * usecase that I can image where you need to set this is if you have
     * a misconfigured reverse proxy that somehow doesn't include the
     * `http-x-forwarded-proto` header into the forwarded requests
     */
    $settings['https'] = strtolower(trim($val)) === 'true';
}

$val = getenv('CHEVERETO_THEME');
if ($val !== false) {
    /*
     * Specify the theme that you wish to use on your Chevereto
     * the default is `Peafowl` - the theme that's included by default with
     * Chevereto free
     * If you specify this environment variable, please make sure to mount
     * the theme directory under /var/www/html/app/themes in the container
     */
    $settings['theme'] = $val;
}
