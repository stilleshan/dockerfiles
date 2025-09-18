import React, { useState, useEffect } from 'react';
import './App.css';
import {
  LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer, CartesianGrid
} from 'recharts';
import { Network, Activity, Calendar, Clock, TrendingUp, Download, Upload, Server, Github } from 'lucide-react';
import { HourlyTable, DailyTable, MonthlyTable, YearlyTable } from './components/TrafficTable';


function formatDate({ year, month, day }) {
  const date = new Date(year, month - 1, day);
  const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  return `${months[date.getMonth()]} ${date.getDate().toString().padStart(2, '0')}, ${date.getFullYear()}`;
}

function formatTime({ hour, minute }) {
  const period = hour >= 12 ? 'PM' : 'AM';
  const displayHour = hour % 12 || 12;
  return `${displayHour.toString().padStart(2, '0')}:${minute.toString().padStart(2, '0')} ${period}`;
}

function formatBytes(bytes) {
  if (!bytes || bytes === 0) return '0 B';
  const k = 1024, dm = 2;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}

function formatMonthYear(year, month) {
  const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  return `${months[month - 1]} ${year}`;
}

const TABS = [
  { id: 'Summary', label: 'Summary', icon: Activity },
  { id: 'Hourly', label: 'Hourly', icon: Clock },
  { id: 'Daily', label: 'Daily', icon: Calendar },
  { id: 'Monthly', label: 'Monthly', icon: Calendar },
  { id: 'Yearly', label: 'Yearly', icon: TrendingUp }
];

const INTERFACES = [
  { id: 'eth0', label: 'Ethernet (eth0)', icon: Network },
  { id: 'wlan0', label: 'Wifi (wlan0)', icon: Network },
  { id: 'docker0', label: 'Docker (docker0)', icon: Server },
  { id: 'tailscale0', label: 'Tailscale (tailscale0)', icon: Network }
];

function App() {
  const [selected, setSelected] = useState('eth0');
  const [tab, setTab] = useState('Summary');
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    fetch(`/api/vnstat/${selected}`)
      .then(res => res.json())
      .then(setData)
      .finally(() => setLoading(false));
  }, [selected]);


  const ifaceInfo = data && data.interfaces ? data.interfaces[0] : null;
  const traffic = ifaceInfo ? ifaceInfo.traffic : null;
  const hourly = traffic && traffic.hour
    ? [...traffic.hour]
      .filter(row => row.time && typeof row.time.hour === 'number')
      .sort((a, b) => {
        const dateA = new Date(a.date.year, a.date.month - 1, a.date.day, a.time.hour);
        const dateB = new Date(b.date.year, b.date.month - 1, b.date.day, b.time.hour);
        return dateB - dateA;
      })
      .slice(0, 24)
    : [];

  const daily = traffic && traffic.day
    ? [...traffic.day].sort((a, b) => {
      const dateA = new Date(a.date.year, a.date.month - 1, a.date.day);
      const dateB = new Date(b.date.year, b.date.month - 1, b.date.day);
      return dateB - dateA;
    })
    : [];

  const monthly = traffic && traffic.month
    ? [...traffic.month].sort((a, b) => {
      const dateA = new Date(a.date.year, a.date.month - 1);
      const dateB = new Date(b.date.year, b.date.month - 1);
      return dateB - dateA;
    })
    : [];

  const yearly = traffic && traffic.year
    ? [...traffic.year].sort((a, b) => b.date.year - a.date.year)
    : [];

  const fivemin = traffic && traffic.fiveminute
    ? traffic.fiveminute.slice(-10).reverse()
    : [];

  const getLabel = (row, type) => {
    if (type === 'hourly') {
      const date = new Date(row.date.year, row.date.month - 1, row.date.day, row.time.hour);
      return date.toLocaleString('en-US', {
        month: 'short',
        day: '2-digit',
        hour: '2-digit',
        hour12: true
      });
    }
    if (type === 'daily') return formatDate(row.date);
    if (type === 'monthly') return formatMonthYear(row.date.year, row.date.month);
    if (type === 'yearly') return `${row.date.year}`;
    return '';
  };

  const graphData = (rows, type) => rows.map(row => ({
    name: getLabel(row, type),
    RX: row.rx ? row.rx : 0,
    TX: row.tx ? row.tx : 0,
    Total: row.rx && row.tx ? row.rx + row.tx : 0,
  }));

  const CustomTooltip = ({ active, payload, label }) => {
    if (active && payload && payload.length) {
      return (
        <div className="bg-gray-800 border border-gray-700 rounded-lg p-3 shadow-xl">
          <p className="text-gray-300 text-sm mb-2">{label}</p>
          {payload.map((entry, index) => (
            <div key={index} className="flex items-center gap-2 text-sm">
              <div
                className="w-3 h-3 rounded-full"
                style={{ backgroundColor: entry.color }}
              />
              <span className="text-gray-300">{entry.dataKey}:</span>
              <span className="text-white font-medium">{formatBytes(entry.value)}</span>
            </div>
          ))}
        </div>
      );
    }
    return null;
  };

  return (
    <div className="min-h-screen bg-gray-950 text-white mb-8">
      <div className="container mx-auto px-4 py-8 max-w-xl w-full">
        {/* Header */}
        <div className="mb-8 text-center">
          <h1 className="text-4xl font-bold mb-2 bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
            Network Traffic Dashboard
          </h1>
          <p className="text-gray-400">Monitor your network interface statistics in real-time</p>
        </div>
        <div class="github">
          <a class="github-icon" href="https://github.com/Kshitiz-b/vnstat-dashboard" target="_blank" rel="noreferrer">
            <Github class="h-5 w-5" />
            <span>Kshitiz-b</span>
          </a>

        </div>

        {/* Interface Selector */}
        <div className="mb-8">
          <label className="block text-sm font-medium text-gray-300 mb-3">
            Select Network Interface
          </label>
          <div className="flex flex-row items-center gap-4 justify-center">
            <Network className="h-5 w-5 text-gray-400" />
            <select
              value={selected}
              onChange={e => setSelected(e.target.value)}
              className="w-full max-w-xs bg-gray-800 border border-gray-700 rounded-lg px-4 py-3 pr-10 text-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all appearance-none"
            >
              {INTERFACES.map(iface => (
                <option key={iface.id} value={iface.id}>
                  {iface.label}
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Tab Navigation */}
        <div className="mb-8">
          <div className="overflow-x-auto no-scrollbar">
            <div className="tab-bar">
              {TABS.map(t => {
                const Icon = t.icon;
                return (
                  <button
                    key={t.id}
                    onClick={() => setTab(t.id)}
                    className={`tab-button${t.id === tab ? ' active' : ''}`}
                  >
                    <Icon className="h-4 w-4" />
                    {t.label}
                  </button>
                );
              })}
            </div>
          </div>
        </div>

        {/* Main Content */}
        <div className="bg-gray-900 rounded-xl border border-gray-800 overflow-hidden">
          {loading ? (
            <div className="flex items-center justify-center h-96">
              <div className="flex items-center gap-3">
                <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
                <span className="text-gray-400">Loading network data...</span>
              </div>
            </div>
          ) : !ifaceInfo ? (
            <div className="flex items-center justify-center h-96">
              <div className="text-center">
                <Server className="h-12 w-12 text-gray-500 mx-auto mb-4" />
                <p className="text-gray-400">No data available for this interface</p>
              </div>
            </div>
          ) : tab === "Summary" ? (
            <div className="p-8">
              <div className="mb-8">
                <h2 className="text-2xl font-bold mb-6 flex items-center gap-3">
                  <Activity className="h-6 w-6 text-blue-400" />
                  {ifaceInfo.name} Overview
                </h2>

                {/* Stats Grid */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
                  <div className="bg-gray-800 rounded-lg p-6 border border-gray-700">
                    <div className="flex items-center gap-3 mb-2">
                      <Download className="h-5 w-5 text-green-400" />
                      <span className="text-sm text-gray-400">Total Received</span>
                    </div>
                    <div className="text-2xl font-bold text-green-400">
                      {formatBytes(traffic.total.rx)}
                    </div>
                  </div>

                  <div className="bg-gray-800 rounded-lg p-6 border border-gray-700">
                    <div className="flex items-center gap-3 mb-2">
                      <Upload className="h-5 w-5 text-blue-400" />
                      <span className="text-sm text-gray-400">Total Sent</span>
                    </div>
                    <div className="text-2xl font-bold text-blue-400">
                      {formatBytes(traffic.total.tx)}
                    </div>
                  </div>

                  <div className="bg-gray-800 rounded-lg p-6 border border-gray-700">
                    <div className="flex items-center gap-3 mb-2">
                      <Calendar className="h-5 w-5 text-purple-400" />
                      <span className="text-sm text-gray-400">Created</span>
                    </div>
                    <div className="text-lg font-semibold text-purple-400">
                      {formatDate(ifaceInfo.created.date)}
                    </div>
                  </div>

                  <div className="bg-gray-800 rounded-lg p-6 border border-gray-700">
                    <div className="flex items-center gap-3 mb-2">
                      <Clock className="h-5 w-5 text-orange-400" />
                      <span className="text-sm text-gray-400">Last Updated</span>
                    </div>
                    <div className="text-lg font-semibold text-orange-400">
                      {formatDate(ifaceInfo.updated.date)}
                    </div>
                    <div className="text-sm text-gray-400">
                      {formatTime(ifaceInfo.updated.time)}
                    </div>
                  </div>
                </div>
              </div>

              {/* Recent Traffic Table */}
              <div>
                <h3 className="text-xl font-semibold mb-4 flex items-center gap-2">
                  <TrendingUp className="h-5 w-5 text-blue-400" />
                  Recent Traffic (5-minute intervals)
                </h3>
                <div className="overflow-x-auto">
                  <table className="w-full border-collapse">
                    <thead>
                      <tr className="bg-gray-800">
                        <th className="text-left p-4 font-medium text-gray-300 border-b border-gray-700">Date</th>
                        <th className="text-left p-4 font-medium text-gray-300 border-b border-gray-700">Time</th>
                        <th className="text-left p-4 font-medium text-gray-300 border-b border-gray-700">
                          <div className="flex items-center gap-2">
                            <Download className="h-4 w-4 text-green-400" />
                            <span className="label-text">Received</span>
                          </div>
                        </th>
                        <th className="text-left p-4 font-medium text-gray-300 border-b border-gray-700">
                          <div className="flex items-center gap-2">
                            <Upload className="h-4 w-4 text-blue-400" />
                            <span className="label-text">Sent</span>
                          </div>
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      {fivemin.map((row, i) => (
                        <tr key={i} className="hover:bg-gray-800 transition-colors">
                          <td className="p-4 border-b border-gray-800 text-gray-300">
                            {formatDate(row.date)}
                          </td>
                          <td className="p-4 border-b border-gray-800 text-gray-300">
                            {formatTime(row.time)}
                          </td>
                          <td className="p-4 border-b border-gray-800 font-medium text-green-400">
                            {formatBytes(row.rx)}
                          </td>
                          <td className="p-4 border-b border-gray-800 font-medium text-blue-400">
                            {formatBytes(row.tx)}
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          ) : (
            <div className="p-8">
              <h2 className="text-2xl font-bold mb-6 flex items-center gap-3">
                <TrendingUp className="h-6 w-6 text-blue-400" />
                {tab} Traffic Analysis
              </h2>
              <div className="bg-gray-800 rounded-lg p-6 border border-gray-700">
                <ResponsiveContainer width="100%" height={400}>
                  <LineChart
                    data={graphData(
                      tab === "Hourly" ? [...hourly.slice(-24)].reverse() :
                        tab === "Daily" ? [...daily].reverse() :
                          tab === "Monthly" ? [...monthly].reverse() :
                            tab === "Yearly" ? [...yearly].reverse() : [],
                      tab.toLowerCase()
                    )}
                    margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
                  >

                    <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
                    <XAxis
                      dataKey="name"
                      stroke="#9CA3AF"
                      fontSize={12}
                      tickLine={false}
                    />
                    <YAxis
                      tickFormatter={formatBytes}
                      stroke="#9CA3AF"
                      fontSize={12}
                      tickLine={false}
                      width={80}
                    />
                    <Tooltip content={<CustomTooltip />} />
                    <Line
                      type="monotone"
                      dataKey="RX"
                      stroke="#10B981"
                      strokeWidth={3}
                      dot={{ fill: '#10B981', strokeWidth: 2, r: 4 }}
                      activeDot={{ r: 6, stroke: '#10B981', strokeWidth: 2 }}
                    />
                    <Line
                      type="monotone"
                      dataKey="TX"
                      stroke="#3B82F6"
                      strokeWidth={3}
                      dot={{ fill: '#3B82F6', strokeWidth: 2, r: 4 }}
                      activeDot={{ r: 6, stroke: '#3B82F6', strokeWidth: 2 }}
                    />
                  </LineChart>
                </ResponsiveContainer>
              </div>

              {/* Summary Stats */}
              {tab === 'Daily' && daily.length > 0 && (
                <div className="mt-6 bg-gray-800 rounded-lg p-6 border border-gray-700">
                  <h4 className="text-lg font-semibold mb-2 text-blue-400">Today's Usage</h4>
                  <div className="flex flex-col sm:flex-row gap-4">
                    <div className="flex flex-col bg-gray-900 rounded-md p-4 border border-gray-700 min-w-[120px] items-center">
                      <span className="text-sm text-gray-400 mb-1">Download:</span>
                      <span className="text-xl font-bold text-green-400 ml-2">{formatBytes(daily[0].rx)}</span>
                    </div>
                    <div className="flex flex-col bg-gray-900 rounded-md p-4 border border-gray-700 min-w-[120px] items-center">
                      <span className="text-sm text-gray-400 mb-1">Upload:</span>
                      <span className="text-xl font-bold text-blue-400 ml-2">{formatBytes(daily[0].tx)}</span>
                    </div>
                  </div>
                </div>
              )}

              {tab === 'Monthly' && monthly.length > 0 && (
                <div className="mt-6 bg-gray-800 rounded-lg p-6 border border-gray-700">
                  <h4 className="text-lg font-semibold mb-2 text-blue-400">This Month's Usage</h4>
                  <div className="flex flex-col sm:flex-row gap-4">
                    <div className="flex flex-col bg-gray-900 rounded-md p-4 border border-gray-700 min-w-[120px] items-center">
                      <span className="text-sm text-gray-400">Download:</span>
                      <span className="text-xl font-bold text-green-400 ml-2">{formatBytes(monthly[0].rx)}</span>
                    </div>
                    <div className="flex flex-col bg-gray-900 rounded-md p-4 border border-gray-700 min-w-[120px] items-center">
                      <span className="text-sm text-gray-400">Upload:</span>
                      <span className="text-xl font-bold text-blue-400 ml-2">{formatBytes(monthly[0].tx)}</span>
                    </div>
                  </div>

                </div>
              )}

              {tab === 'Yearly' && yearly.length > 0 && (
                <div className="mt-6 bg-gray-800 rounded-lg p-6 border border-gray-700">
                  <h4 className="text-lg font-semibold mb-2 text-blue-400">This Year's Usage</h4>
                  <div className="flex flex-col sm:flex-row gap-4">
                    <div className="flex flex-col bg-gray-900 rounded-md p-4 border border-gray-700 min-w-[120px] items-center">
                      <span className="text-sm text-gray-400">Download:</span>
                      <span className="text-xl font-bold text-green-400 ml-2">{formatBytes(yearly[0].rx)}</span>
                    </div>
                    <div className="flex flex-col bg-gray-900 rounded-md p-4 border border-gray-700 min-w-[120px] items-center">
                      <span className="text-sm text-gray-400">Upload:</span>
                      <span className="text-xl font-bold text-blue-400 ml-2">{formatBytes(yearly[0].tx)}</span>
                    </div>
                  </div>
                </div>
              )}
              <div>
                {tab === 'Hourly' && <HourlyTable data={hourly} />}
                {tab === 'Daily' && <DailyTable data={daily} />}
                {tab === 'Monthly' && <MonthlyTable data={monthly} />}
                {tab === 'Yearly' && <YearlyTable data={yearly} />}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;