import "dotenv/config";

export const config = Object.freeze({
	ENV: process.env.ENV || "dev",
	WORKERS: Number.parseInt(process.env.WORKERS || "1"),

	RMQ_HOSTNAME: process.env.RMQ_HOSTNAME || "",
	RMQ_USER: process.env.RMQ_USER || "",
	RMQ_PASSWORD: process.env.RMQ_PASSWORD || "",
	RMQ_VHOST: process.env.RMQ_VHOST || "",
	RMQ_QUEUE: process.env.RMQ_QUEUE || "",

	TARGET_BUCKET_SIZE: Number.parseInt(process.env.TARGET_BUCKET_SIZE || "1000"),
	TARGET_RESTORE_RATE: Number.parseInt(
		process.env.TARGET_RESTORE_RATE || "100",
	),
	TARGET_INVOCATION_COST: Number.parseInt(
		process.env.TARGET_INVOCATION_COST || "50",
	),
});
