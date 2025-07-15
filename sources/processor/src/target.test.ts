import target from "./target";

describe("target.getWaitBeforeInvocation()", () => {
	const mockedTimestamp = 1752539900000;
	// Configure target to testable values
	target.bucketRestoreRate = 10;
	target.invocationCost = 20;
	target.bucketSize = 1000;

	// Set current time to 5 seconds since last update
	jest.useFakeTimers().setSystemTime(mockedTimestamp);

	// The current bucket has to be 5 * 10 = 50 > 20 (invocation cost)
	it("should return 0 to allow immediate invocation", () => {
		target.lastBucketUpdate = mockedTimestamp - 5000;
		target.currentBucket = 0;
		expect(target.getWaitBeforeInvocation()).toBe(0);
	});

	// The current bucket has to be 1 * 10 = 10 < 20 (invocation cost)
	// Wait period should be time required to reach 2.5 * 20 = 50
	// Wait therfore has to be an additional 4 seconds
	it("should return 4 seconds of wait", () => {
		target.lastBucketUpdate = mockedTimestamp - 1000;
		target.currentBucket = 0;
		expect(target.getWaitBeforeInvocation()).toBe(4000);
	});
});
