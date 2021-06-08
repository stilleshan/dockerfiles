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

$logk = log(1024);

function getScale($bytes)
{
    global $logk;

    $ui = floor(round(log($bytes)/$logk,3));
    if ($ui < 0) { $ui = 0; }
    if ($ui > 8) { $ui = 8; }

    return $ui;
}

// Get the largest value in an array
function getLargestValue($array) {
    return $max = array_reduce($array, function ($a, $b) {
        return $a > $b['total'] ? $a : $b['total'];
    });
}

function getBaseValue($array, $scale)
{
    $big = pow(1024,9);

    // Find the smallest non-zero value
    $sml = array_reduce($array, function ($a, $b) {
        if  ((1 <= $b['rx']) && ($b['rx'] < $b['tx'])) {
            $sm = $b['rx'];
        } else {
            $sm = $b['tx'];
        }
        if (($sm < 1) || ($a < $sm)) {
            return $a;
        } else {
            return $sm;
        }
    }, $big);

    if ($sml >= $big/2) {
        $sml = 1;
    }

    // divide by scale then round down to a power of 10
    $base = pow(10,floor(round(log10($sml/pow(1024,$scale)),3)));

    // convert back to bytes
    $baseByte = $base * pow(1024, $scale);

    // Don't make the bar invisable - must be > 5% difference
    if ($sml / $baseByte < 1.05) {
        $base = $base / 10;
    }

    return $base;
}

function formatSize($bytes, $vnstatJsonVersion, $decimals = 2) {

    // json version 1 = convert from KiB
    // json version 2 = convert from bytes
    if ($vnstatJsonVersion == 1) {
        $bytes *= 1024;  // convert from kibibytes to bytes
    }

    return formatBytes($bytes, $decimals);
}

function getLargestPrefix($scale)
{
    $suffixes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    return $suffixes[$scale];
}

function formatBytes($bytes, $decimals = 3) {

    $scale = getScale($bytes);

    return round($bytes/pow(1024, $scale), $decimals) .' '. getLargestPrefix($scale);
}

function formatBytesTo($bytes, $scale, $decimals = 4) {

    if ($bytes == 0) {
        return '0';
    }

    return number_format(($bytes / pow(1024, $scale)), $decimals, ".", "");
}

function kibibytesToBytes($kibibytes, $vnstatJsonVersion) {
    if ($vnstatJsonVersion == 1) {
        return $kibibytes *= 1024;
    } else {
        return $kibibytes;
    }
}

function sortingFunction($item1, $item2) {
    if ($item1['time'] == $item2['time']) {
        return 0;
    } else {
        return $item1['time'] > $item2['time'] ? -1 : 1;
    }
};

?>
