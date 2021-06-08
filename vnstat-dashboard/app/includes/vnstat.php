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

class vnStat {
	protected $executablePath;
	protected $vnstatVersion;
	protected $vnstatJsonVersion;
	protected $vnstatData;

	public function __construct ($executablePath) {
		if (isset($executablePath)) {
			$this->executablePath = $executablePath;

			// Execute a command to output a json dump of the vnstat data
			$vnstatStream = popen("$this->executablePath --json", 'r');

			// Is the stream valid?
			if (is_resource($vnstatStream)) {
				$streamBuffer = '';

				while (!feof($vnstatStream)) {
					$streamBuffer .= fgets($vnstatStream);
				}

				// Close the handle
				pclose($vnstatStream);

				$this->processVnstatData($streamBuffer);
			} else {

			}


		} else {
			die();
		}
	}

	private function processVnstatData($vnstatJson) {
		$decodedJson = json_decode($vnstatJson, true);

		// Check the JSON is valid
		if (json_last_error() != JSON_ERROR_NONE) {
			throw new Exception('JSON is invalid');
		}

		$this->vnstatData = $decodedJson;
		$this->vnstatVersion = $decodedJson['vnstatversion'];
		$this->vnstatJsonVersion = $decodedJson['jsonversion'];
	}

	public function getVnstatVersion() {
		return $this->vnstatVersion;
	}

	public function getVnstatJsonVersion() {
		return $this->vnstatJsonVersion;
	}

	public function getInterfaces() {
		// Create a placeholder array
		$vnstatInterfaces = [];

		foreach($this->vnstatData['interfaces'] as $interface) {
		    if ($this->vnstatJsonVersion == 1) {
			array_push($vnstatInterfaces, $interface['id']);
		    } else {
			array_push($vnstatInterfaces, $interface['name']);
		    }
		}

		return $vnstatInterfaces;
	}

	public function getInterfaceData($timeperiod, $type, $interface) {
		// If json version equals 1, add an 's' onto the end of each type.
		// e.g. 'top' becomes 'tops'
		$typeAppend = '';
		if ($this->vnstatJsonVersion == 1) {
			$typeAppend = 's';
		}

		// Blank placeholder
		$trafficData = [];
		$i = -1;

		// Get the array index for the chosen interface
		if ($this->vnstatJsonVersion == 1) {
		    $arrayIndex = array_search($interface, array_column($this->vnstatData['interfaces'], 'id'));
                } else {
		    $arrayIndex = array_search($interface, array_column($this->vnstatData['interfaces'], 'name'));
                }
 
		if ($timeperiod == 'top10') {
			if ($type == 'table') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['top'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

						$trafficData[$i]['label'] = date('d/m/Y', strtotime($traffic['date']['month'] . "/" . $traffic['date']['day'] . "/" . $traffic['date']['year']));;
						$trafficData[$i]['rx'] = formatSize($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = formatSize($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = formatSize(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
                                                $trafficData[$i]['totalraw'] = ($traffic['rx'] + $traffic['tx']);
					}
				}
			}
		}

		if (($this->vnstatJsonVersion > 1) && ($timeperiod == 'five')) {
			if ($type == 'table') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['fiveminute'] as $traffic) {
					if (is_array($traffic)) {
						$i++;

						$trafficData[$i]['label'] = date("d/m/Y H:i", mktime($traffic['time']['hour'], $traffic['time']['minute'], 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']));
                                                $trafficData[$i]['time'] =  mktime($traffic['time']['hour'], $traffic['time']['minute'], 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']);
						$trafficData[$i]['rx'] = formatSize($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = formatSize($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = formatSize(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			} else if ($type == 'graph') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['fiveminute'] as $traffic) {
					if (is_array($traffic)) {
						$i++;

						$trafficData[$i]['label'] = sprintf("Date(%d, %d, %d, %d, %d)", $traffic['date']['year'], $traffic['date']['month']-1, $traffic['date']['day'], $traffic['time']['hour'], $traffic['time']['minute']);
                                                $trafficData[$i]['time'] =  mktime($traffic['time']['hour'], $traffic['time']['minute'], 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']);
						$trafficData[$i]['rx'] = kibibytesToBytes($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = kibibytesToBytes($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = kibibytesToBytes(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			}
		}

		if ($timeperiod == 'hourly') {
			if ($type == 'table') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['hour'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

                                                if ($this->vnstatJsonVersion == 1) {
                                                    $hour = $traffic['id'];
                                                } else {
                                                    $hour = $traffic['time']['hour'];
                                                }

						$trafficData[$i]['label'] = date("d/m/Y H:i", mktime($hour, 0, 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']));
                                                $trafficData[$i]['time'] =  mktime($hour, 0, 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']);
						$trafficData[$i]['rx'] = formatSize($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = formatSize($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = formatSize(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}


			} else if ($type == 'graph') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['hour'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

                                                if ($this->vnstatJsonVersion == 1) {
                                                    $hour = $traffic['id'];
                                                } else {
                                                    $hour = $traffic['time']['hour'];
                                                }

						$trafficData[$i]['label'] = sprintf("Date(%d, %d, %d, %d)", $traffic['date']['year'], $traffic['date']['month']-1, $traffic['date']['day'], $hour);
                                                $trafficData[$i]['time'] =  mktime($hour, 0, 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']);
						$trafficData[$i]['rx'] = kibibytesToBytes($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = kibibytesToBytes($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = kibibytesToBytes(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			}
		}

		if ($timeperiod == 'daily') {
			if ($type == 'table') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['day'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

						$trafficData[$i]['label'] = date('d/m/Y', mktime(0, 0, 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']));
                                                $trafficData[$i]['time'] =  mktime(0, 0, 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']);
						$trafficData[$i]['rx'] = formatSize($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = formatSize($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = formatSize(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			} else if ($type == 'graph') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['day'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

						$trafficData[$i]['label'] = sprintf("Date(%d, %d, %d)", $traffic['date']['year'], $traffic['date']['month']-1, $traffic['date']['day']);
                                                $trafficData[$i]['time'] =  mktime(0, 0, 0, $traffic['date']['month'], $traffic['date']['day'], $traffic['date']['year']);
						$trafficData[$i]['rx'] = kibibytesToBytes($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = kibibytesToBytes($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = kibibytesToBytes(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			}
		}

		if ($timeperiod == 'monthly') {
			if ($type == 'table') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['month'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

						$trafficData[$i]['label'] = date('F Y', mktime(0, 0, 0, $traffic['date']['month'], 10, $traffic['date']['year']));
                                                $trafficData[$i]['time'] =  mktime(0, 0, 0, $traffic['date']['month'], 10, $traffic['date']['year']);
						$trafficData[$i]['rx'] = formatSize($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = formatSize($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = formatSize(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			} else if ($type == 'graph') {
				foreach ($this->vnstatData['interfaces'][$arrayIndex]['traffic']['month'.$typeAppend] as $traffic) {
					if (is_array($traffic)) {
						$i++;

                                                $trafficData[$i]['label'] = sprintf("Date(%d, %d, %d)", $traffic['date']['year'], $traffic['date']['month'] - 1, 10);
                                                $trafficData[$i]['time'] =  mktime(0, 0, 0, $traffic['date']['month'], 10, $traffic['date']['year']);
						$trafficData[$i]['rx'] = kibibytesToBytes($traffic['rx'], $this->vnstatJsonVersion);
						$trafficData[$i]['tx'] = kibibytesToBytes($traffic['tx'], $this->vnstatJsonVersion);
						$trafficData[$i]['total'] = kibibytesToBytes(($traffic['rx'] + $traffic['tx']), $this->vnstatJsonVersion);
					}
				}
			}
		}

		if ($timeperiod != 'top10') {
                    usort($trafficData, 'sortingFunction');
                }

                if ($type == 'graph') {
                    // Get the largest value and then prefix (B, KB, MB, GB, etc)
                    $trafficLargestValue = getLargestValue($trafficData);
                    $trafficScale = getScale($trafficLargestValue);
                    $trafficLargestPrefix = getLargestPrefix($trafficScale);
                    $trafficBase = getBaseValue($trafficData, $trafficScale);
                    if (($trafficBase < .0099) && ($trafficScale >= 1))
                    {
                        $trafficScale = $trafficScale - 1;
                        $trafficLargestPrefix = getLargestPrefix($trafficScale);
                        $trafficBase = getBaseValue($trafficData, $trafficScale);
                    }

                    foreach($trafficData as &$value) {
                        $value['rx'] = formatBytesTo($value['rx'], $trafficScale);
                        $value['tx'] = formatBytesTo($value['tx'], $trafficScale);
                        $value['total'] = formatBytesTo($value['total'], $trafficScale);
                    }

                    unset($value);
                    $trafficData[0]['delimiter'] = $trafficLargestPrefix;
                    $trafficData[0]['base'] = $trafficBase;
                }

		return $trafficData;
	}
}

?>
