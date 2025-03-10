// src/lib/utils/auth.ts
import { writable } from 'svelte/store';

export const isAuthenticated = writable(false);

export function login() {
  // Implement login logic
  isAuthenticated.set(true);
}

export function logout() {
  // Implement logout logic
  isAuthenticated.set(false);
}