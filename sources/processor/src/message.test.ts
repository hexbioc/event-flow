import message from "./message";

describe("getMessage()", () => {
	it("should return the correct message when called", () => {
		expect(message.get()).toBe("Hello there!");
	});
});
