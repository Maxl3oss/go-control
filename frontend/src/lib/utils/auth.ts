// src/lib/utils/auth.ts
import { jwtDecode } from 'jwt-decode';
import { writable } from 'svelte/store';
import { BASE_URL } from './helper';

type IUserAuth = {
  isAuthenticated: boolean
  token: string
  email: string
  expiresAt: number | null
}

export const userAuth = writable<IUserAuth>({
  email: "",
  token: "",
  isAuthenticated: false,
  expiresAt: null,
});

export async function login(email: string, password: string) {
  // Implement login logic
  const response = await fetch(`${BASE_URL}/login`, {
    method: 'POST',
    body: JSON.stringify({ email, password })
  });
  const data = await response.json();
  if (response.ok) {
    const decodedToken = jwtDecode(data?.token);
    const expiresAt = (decodedToken?.exp || 0) * 1000; 
    userAuth.set({
      email: data?.email ?? "",
      token: data?.token ?? "",
      isAuthenticated: true,
      expiresAt: expiresAt,
    });
    localStorage.setItem("expiresAt", String(expiresAt));
    localStorage.setItem("token", data?.token ?? "")
    return data;
  } else {
    logout();
    return data;
  }
}

export function logout() {
  // Implement logout logic
  localStorage.removeItem("token");
  localStorage.removeItem("expiresAt");
  userAuth.set({
    email: "",
    token: "",
    isAuthenticated: false,
    expiresAt: null
  });
}
