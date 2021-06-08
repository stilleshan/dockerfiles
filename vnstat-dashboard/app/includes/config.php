<?php

/*
 * Copyright (C) 2019 Alexander Marston (alexander.marston@gmail.com)
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Uncomment to enable error reporting to the screen
/*ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
ferror_reporting(E_ALL);*/

// Set the default system Timezone
date_default_timezone_set('Europe/London');

// Path of vnstat
$vnstat_bin_dir = '/usr/bin/vnstat';

// Path of config file
/*$vnstat_config = '/etc/vnstat.conf';*/

// linear or logarithmic graphs. Uncomment for logarithmic
/*$graph_type = 'log';*/

// Set to true to set your own interfaces
$use_predefined_interfaces = false;

if ($use_predefined_interfaces == true) {
    $interface_list = ["eth0", "eth1"];

    $interface_name['eth0'] = "Internal #1";
    $interface_name['eth1'] = "Internal #2";
}
