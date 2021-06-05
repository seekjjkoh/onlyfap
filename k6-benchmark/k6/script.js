import http from "k6/http";

const URL = __ENV.URL

export let options = {
  summaryTrendStats: ["avg", "min", "med", "max", "p(95)", "p(99)", "p(99.99)", "count"],
};

export function setup() {
  return {
    url: URL,
  };
}

export default function(data) {
  http.get(data.url);
}

export function teardown(data) { }
