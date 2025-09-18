import React from 'react';
import { Clock, Calendar, Download, Upload, BarChart, LineChart } from 'lucide-react';

const TrafficTable = ({ title, icon, data, headers }) => (
  <div>
    <h3 className="text-xl font-semibold mb-4 flex items-center gap-2">
      {icon}
      {title}
    </h3>
    <div className="overflow-x-auto">
      <table className="w-full border-collapse">
        <thead>
          <tr className="bg-gray-800">
            {headers.map((header, i) => (
              <th key={i} className="text-left p-4 font-medium text-gray-300 border-b border-gray-700">
                <div className="flex items-center gap-2">
                  {header.icon}
                  <span className="label-text">{header.label}</span>
                </div>
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((row, i) => (
            <tr key={i} className="hover:bg-gray-800 transition-colors">
              {headers.map((header, j) => (
                <td
                  key={j}
                  className={`p-4 border-b border-gray-800 text-gray-300 ${header.className || ''}`.trim()}
                >
                  {header.render(row)}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  </div>
);

const formatBytes = (bytes) => {
  if (!bytes || bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
};

const HourlyTable = ({ data }) => {
  const sortedData = [...data]
    .filter((row) => row.rx || row.tx)
    .sort((a, b) => {
      const dateA = new Date(a.date.year, a.date.month - 1, a.date.day, a.hour || 0);
      const dateB = new Date(b.date.year, b.date.month - 1, b.date.day, b.hour || 0);
      return dateB - dateA; // descending
    })
    .slice(0, 24); // last 24 hours

  return (
    <TrafficTable
      title="Hourly Traffic"
      icon={<Clock className="h-5 w-5 text-yellow-400" />}
      data={sortedData}
      headers={[
        {
          label: 'Date',
          render: (row) =>
            new Date(row.date.year, row.date.month - 1, row.date.day).toLocaleDateString('en-US', {
              year: 'numeric',
              month: 'short',
              day: '2-digit',
            }),
        },
        {
          label: 'Time',
          render: (row) =>
            new Date(
              row.date.year,
              row.date.month - 1,
              row.date.day,
              row.time?.hour ?? 0,
              row.time?.minute ?? 0
            ).toLocaleTimeString('en-US', {
              hour: '2-digit',
              minute: '2-digit',
              hour12: true,
            }),
        },
        {
          label: 'Received',
          icon: <Download className="h-4 w-4 text-green-400" />,
          className: 'font-medium text-green-400',
          render: (row) => formatBytes(row.rx),
        },
        {
          label: 'Sent',
          icon: <Upload className="h-4 w-4 text-blue-400" />,
          className: 'font-medium text-blue-400',
          render: (row) => formatBytes(row.tx),
        },
      ]}
    />
  );
};

const DailyTable = ({ data }) => (
  <TrafficTable
    title="Daily Traffic"
    icon={<Calendar className="h-5 w-5 text-green-400" />}
    data={data}
    headers={[
      { label: 'Date',
          render: (row) =>
            new Date(row.date.year, row.date.month - 1, row.date.day).toLocaleDateString('en-US', {
              year: 'numeric',
              month: 'short',
              day: '2-digit',
            }), },
      {
        label: 'Received',
        icon: <Download className="h-4 w-4 text-green-400" />,
        className: 'font-medium text-green-400',
        render: (row) => formatBytes(row.rx),
      },
      {
        label: 'Sent',
        icon: <Upload className="h-4 w-4 text-blue-400" />,
        className: 'font-medium text-blue-400',
        render: (row) => formatBytes(row.tx),
      },
    ]}
  />
);

const MonthlyTable = ({ data }) => (
  <TrafficTable
    title="Monthly Traffic"
    icon={<BarChart className="h-5 w-5 text-purple-400" />}
    data={data}
    headers={[
      { label: 'Month', render: (row) =>
            new Date(row.date.year, row.date.month - 1).toLocaleDateString('en-US', {
              year: 'numeric',
              month: 'short',
            }), },
      {
        label: 'Received',
        icon: <Download className="h-4 w-4 text-green-400" />,
        className: 'font-medium text-green-400',
        render: (row) => formatBytes(row.rx),
      },
      {
        label: 'Sent',
        icon: <Upload className="h-4 w-4 text-blue-400" />,
        className: 'font-medium text-blue-400',
        render: (row) => formatBytes(row.tx),
      },
    ]}
  />
);

const YearlyTable = ({ data }) => (
  <TrafficTable
    title="Yearly Traffic"
    icon={<LineChart className="h-5 w-5 text-orange-400" />}
    data={data}
    headers={[
      { label: 'Year', render: (row) => `${row.date.year}` },
      {
        label: 'Received',
        icon: <Download className="h-4 w-4 text-green-400" />,
        className: 'font-medium text-green-400',
        render: (row) => formatBytes(row.rx),
      },
      {
        label: 'Sent',
        icon: <Upload className="h-4 w-4 text-blue-400" />,
        className: 'font-medium text-blue-400',
        render: (row) => formatBytes(row.tx),
      },
    ]}
  />
);

export { HourlyTable, DailyTable, MonthlyTable, YearlyTable };
