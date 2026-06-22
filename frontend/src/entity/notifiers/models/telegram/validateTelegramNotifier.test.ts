import { describe, expect, it } from 'vitest';

import { validateTelegramNotifier } from './validateTelegramNotifier';

describe('validateTelegramNotifier', () => {
  it('requires proxy URL when HTTP proxy is enabled for a new notifier', () => {
    expect(
      validateTelegramNotifier(true, {
        botToken: 'token',
        targetChatId: '123456',
        isHttpProxyEnabled: true,
      }),
    ).toBe(false);
  });

  it('allows an existing hidden proxy URL', () => {
    expect(
      validateTelegramNotifier(false, {
        botToken: '',
        targetChatId: '123456',
        isHttpProxyEnabled: true,
        hasHttpProxyUrl: true,
      }),
    ).toBe(true);
  });

  it('rejects non-http proxy URLs', () => {
    expect(
      validateTelegramNotifier(true, {
        botToken: 'token',
        targetChatId: '123456',
        isHttpProxyEnabled: true,
        httpProxyUrl: 'https://proxy.example.com:3128',
      }),
    ).toBe(false);
  });

  it('allows valid http proxy URLs', () => {
    expect(
      validateTelegramNotifier(true, {
        botToken: 'token',
        targetChatId: '123456',
        isHttpProxyEnabled: true,
        httpProxyUrl: 'http://user:password@proxy.example.com:3128',
      }),
    ).toBe(true);
  });
});
