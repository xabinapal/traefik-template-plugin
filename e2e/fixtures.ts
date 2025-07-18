import {
  test as base,
  request as baseRequest,
  APIRequestContext,
} from "@playwright/test";

interface TestFixtures {
  request: APIRequestContext;
}

interface WorkerFixtures {
}

export const test = base.extend<TestFixtures, WorkerFixtures>({
  request: async ({}, use) => {
    const context = await baseRequest.newContext();
    const originalFetch = context.fetch.bind(context);

    context.fetch = async (url, options = {}) => {
      const originalUrl: URL =
        typeof url === "string" ? new URL(url) : new URL(url.url());

      const rewrittenUrl = new URL(originalUrl.href);
      rewrittenUrl.host = "localhost";

      options.headers = {
        ...options.headers,
        Host: originalUrl.host,
      };

      options.maxRedirects = 0;

      return originalFetch(rewrittenUrl.href, options);
    };

    await use(context);
    await context.dispose();
  },
});
