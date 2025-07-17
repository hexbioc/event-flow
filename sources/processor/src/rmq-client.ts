import { config } from "./config";
import amqp from "amqplib";
class RabbitMQClient {
	uri: string;

	connection: amqp.ChannelModel | undefined;

	constructor() {
		const proto = config.RMQ_TLS === "true" ? "amqps" : "amqp";
		this.uri =
			`${proto}://${config.RMQ_USER}:${config.RMQ_PASSWORD}` +
			`@${config.RMQ_HOSTNAME}/${config.RMQ_VHOST}`;
	}

	async connect() {
		// Close any previous connections
		await this.close();

		this.connection = await amqp.connect(this.uri);
	}

	async assertQueue(queueName: string, opts?: amqp.Options.AssertQueue) {
		if (!this.connection) {
			throw new Error("Connection not yet established");
		}

		const channel = await this.connection?.createChannel();
		await channel.assertQueue(queueName, opts);

		await channel.close();
	}

	async close() {
		if (this.connection) {
			try {
				await this.connection.close();
				this.connection = undefined;
			} catch {}
		}
	}

	async newChannel() {
		if (!this.connection) {
			throw new Error("Connection not yet established");
		}
		return await this.connection.createChannel();
	}
}

export default new RabbitMQClient();
