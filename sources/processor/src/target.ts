import { config } from "./config";
import type { TargetPayloadDTO } from "./dto";
import { loggerForModule } from "./logger";
import { sleep } from "./utils";

// This indicates the period to wait before invoking again
const RestorationWaitFactor = 2.5;

class TargetManager {
	logger = loggerForModule("target");
	bucketSize: number;
	bucketRestoreRate: number;
	invocationCost: number;

	currentBucket: number;
	lastBucketUpdate: number;

	constructor() {
		this.bucketSize = config.TARGET_BUCKET_SIZE;
		this.bucketRestoreRate = config.TARGET_RESTORE_RATE;
		this.invocationCost = config.TARGET_INVOCATION_COST;

		this.currentBucket = this.bucketSize;
		this.lastBucketUpdate = Date.now();
	}

	getWaitBeforeInvocation() {
		const currentTimestamp = Date.now();

		// Get updated bucket size
		const updatedBucket = Math.min(
			this.bucketSize,
			this.currentBucket +
				Math.floor((currentTimestamp - this.lastBucketUpdate) / 1000) *
					this.bucketRestoreRate,
		);

		if (this.invocationCost < updatedBucket) {
			this.currentBucket = updatedBucket - this.invocationCost;
			this.lastBucketUpdate = Date.now();

			return 0;
		}

		// Bucket does not have sufficient tokens
		return Math.ceil(
			1000 *
				((this.invocationCost * RestorationWaitFactor - updatedBucket) /
					this.bucketRestoreRate),
		);
	}

	async invoke(payload: TargetPayloadDTO) {
		this.logger.info(`Target invoked with payload: ${JSON.stringify(payload)}`);

		const waitPeriodMillis = this.getWaitBeforeInvocation();

		if (waitPeriodMillis === 0) {
			// Emulate processing
			await sleep(Math.random() * 1000);

			return;
		}

		this.logger.warn(
			`Insufficient tokens in bucket, will attempt after ${(waitPeriodMillis / 1000).toFixed(2)} seconds`,
		);
		await sleep(waitPeriodMillis);
		await this.invoke(payload);
	}
}

export default new TargetManager();
