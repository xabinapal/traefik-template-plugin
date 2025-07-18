import path from "path";

import { Page, chromium } from "@playwright/test";

import * as dockerCompose from "docker-compose";

export default async function globalSetup() {
  console.log("Starting docker compose sandbox services...");

  await dockerCompose.upAll({
    cwd: path.join(__dirname, "../sandbox"),
    commandOptions: ["--wait"],
  });

  console.log("Started Docker compose sandbox services");

  console.log("Waiting for docker compose sandbox services to be ready...");

  const browser = await chromium.launch();
  const page = await browser.newPage();

  await waitForTraefik(page);

  await browser.close();

  console.log("Docker compose sandbox services are ready");
}

async function waitForTraefik(page: Page) {
  let traefikStatusCode = 0;
  while (traefikStatusCode !== 200) {
    const traefik = await page.goto("http://traefik.localhost:8080/ping/");
    traefikStatusCode = traefik?.status() || 0;

    if (traefikStatusCode !== 200) {
      console.log("Traefik not ready yet...");
      await new Promise((resolve) => setTimeout(resolve, 1000));
    }
  }
}
