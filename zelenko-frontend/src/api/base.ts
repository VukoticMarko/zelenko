const BASE_URL = "http://localhost:2023";

export const fetchLike = async (method: string, url: string, body?: any) => {
    const response = await fetch(`${BASE_URL}/${url}`, {
        method,
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body)
      });
      return response.json();
}