export async function fetchWithAuth(
  url: string,
  options: RequestInit = {},
  onInvalidToken: () => void
) {
  const res = await fetch(url, options);

  if (res.status === 401) {
    try {
      const data = await res.json();
      if (data?.error === "Token inválido") {
        onInvalidToken();
        return null;
      }
    } catch {
      onInvalidToken();
      return null;
    }
  }

  return res;
}
