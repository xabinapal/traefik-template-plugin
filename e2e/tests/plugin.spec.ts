import { expect } from "@playwright/test";
import { StatusCodes } from "http-status-codes";

import { test } from "../fixtures";

test.describe("plugin", () => {
  test("should add headers to the request", async ({ request }) => {
    const response = await request.get("http://whoami.localhost/");
    expect(response.status()).toBe(StatusCodes.OK);

    const body = await response.text();
    expect(body).toContain("X-Demo: test");
    expect(body).toContain("X-Url: http://localhost");
  });
});
