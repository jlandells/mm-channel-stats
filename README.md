# mm-channel-stats

`mm-channel-stats` is a command-line utility for collecting detailed statistics about Mattermost channels, such as message counts, last post dates, and metadata. The output can be generated in either JSON or CSV format, making it ideal for Mattermost administrators who need insights into channel usage and activity.

---

## Features

- Collects key channel metrics:
  - Message counts (total and root-only)
  - Last post date and last update date
  - Channel type, header, and purpose information
- Supports JSON and CSV output formats
- Configurable via command-line flags, environment variables, and an optional configuration file

---

## Getting Started

### **Download the Binary**

Visit the [Releases](https://github.com/jlandells/mm-channel-stats/releases) page and download the appropriate binary for your operating system.

---

### **Usage**

Run the utility with the required parameters. At a minimum, you must specify the Mattermost instance URL and an API token.

#### **Basic Example**
```bash
./mm-channel-stats -url mattermost.example.com -token YOUR_API_TOKEN
```

This will generate a JSON file named `channel_stats.json` in the current directory.

#### **Generate a CSV File**
```bash
./mm-channel-stats -url mattermost.example.com -token YOUR_API_TOKEN -csv -file=channel_stats.csv
```

This will generate a CSV file named `channel_stats.csv` in the current directory.

---

### **Command-Line Options**

| **Option**       | **Env Var Alternative** | **Required?** | **Description**                                     | **Default**          |
|-------------------|-------------------------|---------------|-----------------------------------------------------|----------------------|
| `-url`           | `MM_URL`                | Yes           | Mattermost instance URL                             |                      |
| `-token`         | `MM_TOKEN`              | Yes           | API token for Mattermost                           |                      |
| `-port`          | `MM_PORT`               | No            | Mattermost port                                    | `443`                |
| `-scheme`        | `MM_SCHEME`             | No            | HTTP scheme (`http` or `https`)                    | `https`              |
| `-csv`           |                         | No            | Generate a CSV output instead of JSON              | `false`              |
| `-file`          |                         | No            | Output filename                                    | `channel_stats.json` or `channel_stats.csv` |
| `-config`        | `MM_CONFIG`             | No            | Path to an optional configuration file             | `./config.json`      |
| `-debug`         | `MM_DEBUG`              | No            | Enable debug mode for additional output            | `false`              |
| `-version`       |                         | No            | Displays the utility version and exits             | `false`              |
| `-help`          |                         | No            | Displays usage information                         | `false`              |

---

### **Configuration**

`mm-channel-stats` supports environment variables and an optional JSON configuration file to simplify repetitive tasks.

#### **Example JSON Config File**
```json
{
  "url": "https://mattermost.example.com",
  "token": "your-mattermost-api-token",
  "csv": true,
  "file": "output.csv",
  "debug": true
}
```

Run the utility with the `-config` flag to use the file:
```bash
./mm-channel-stats -config ./config.json
```

---

## Contributing

We welcome contributions from the community! Whether it's a bug report, a feature suggestion, or a pull request, your input is valuable to us. Please feel free to contribute in the following ways:
- **Issues and Pull Requests**: For specific questions, issues, or suggestions for improvements, open an issue or a pull request in this repository.
- **Mattermost Community**: Join the discussion in the [Integrations and Apps](https://community.mattermost.com/core/channels/integrations) channel on the Mattermost Community server.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions, feedback, or contributions regarding this project, please use the following methods:
- **Issues and Pull Requests**: For specific questions, issues, or suggestions for improvements, feel free to open an issue or a pull request in this repository.
- **Mattermost Community**: Join us in the Mattermost Community server, where we discuss all things related to extending Mattermost. You can find me in the channel [Integrations and Apps](https://community.mattermost.com/core/channels/integrations).
- **Social Media**: Follow and message me on Twitter, where I'm [@jlandells](https://twitter.com/jlandells).
