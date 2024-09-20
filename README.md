
![Logo](https://ik.imagekit.io/pj3r6oe9k/prasorganic-high-resolution-logo-transparent.svg?updatedAt=1726835541390)

# Prasorganic Notification Service

Prasorganic Notification Service is one of the components in the Prasorganic architecture Microservices built with Go (Golang). This service is used to manage notifications from third parties via RESTful API and Message Broker. Kafka is selected as the message broker due to its ability to store messages in files, ensuring important data such as transactions remaining safe and available, even if message management failures occur in the other services.

## Tech Stack

[![My Skills](https://skillicons.dev/icons?i=go,docker,kafka,bash,git&theme=light)](https://skillicons.dev)

## Features

- **Notification Management:** Supports operations for managing notifications from Shipper and Midtrans.

- **RESTful API:** Provides a RESTful API using Fiber with various middleware for managing requests.

- **Message Broker:** This service acts as a producer for the Kafka Notification Service.

- **Logging:** Logs are recorded using Logrus.

- **Error Handling:** Implements error handling to ensure proper detection and handling of errors, minimizing the impact on both the client and server.

- **Configuration and Security:** Utilizes Viper and HashiCorp Vault for integrated configuration and security management.

- **Testing:** Implements unit testing using Testify.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

This project makes use of third-party packages and tools. The licenses for these
dependencies can be found in the `LICENSES` directory.

## Dependencies and Their Licenses


- `Go:` Licensed under the BSD 3-Clause "New" or "Revised" License. For more information, see the [Go License](https://github.com/golang/go/blob/master/LICENSE).

- `Docker:` Licensed under the Apache License 2.0. For more information, see the [Docker License](https://github.com/docker/docs/blob/main/LICENSE).

- `Docker Compose:` Licensed under the Apache License 2.0. For more information, see the [Docker Compose License](https://github.com/docker/compose/blob/main/LICENSE).

- `Kafka:` Licensed under the Apache License 2.0. For more information, see the [Kafka License](https://github.com/apache/kafka/blob/trunk/LICENSE).

- `Bitnami/Kafka:` Licensed under the Apache License 2.0. For more information, see the [Bitnami/Kafka License](https://www.apache.org/licenses/LICENSE-2.0).

- `Zookeeper:` Licensed under the Apache License 2.0. For more information, see the [Zookeper License](https://github.com/apache/zookeeper/blob/master/LICENSE.txt).

- `Bitnami/Zookeeper:` Licensed under the Apache License 2.0. For more information, see the [Bitnami/Zookeper License](https://www.apache.org/licenses/LICENSE-2.0).

- `GNU Make:` Licensed under the GNU General Public License v3.0. For more information, see the [GNU Make License](https://www.gnu.org/licenses/gpl.html).

- `GNU Bash:` Licensed under the GNU General Public License v3.0. For more information, see the [Bash License](https://www.gnu.org/licenses/gpl-3.0.html).

- `Git:` Licensed under the GNU General Public License version 2.0. For more information, see the [Git License](https://opensource.org/license/GPL-2.0).

## Thanks 👍
Thank you for viewing my project.