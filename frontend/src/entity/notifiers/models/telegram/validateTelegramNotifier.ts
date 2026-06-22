import type { TelegramNotifier } from './TelegramNotifier';

export const validateTelegramNotifier = (
  isCreate: boolean,
  notifier: TelegramNotifier,
): boolean => {
  if (isCreate && !notifier.botToken) {
    return false;
  }

  if (!notifier.targetChatId) {
    return false;
  }

  // If thread is enabled, thread ID must be present and valid
  if (notifier.isSendToThreadEnabled && (!notifier.threadId || notifier.threadId <= 0)) {
    return false;
  }

  if (notifier.isHttpProxyEnabled) {
    if (!notifier.httpProxyUrl && !notifier.hasHttpProxyUrl) {
      return false;
    }

    if (notifier.httpProxyUrl) {
      try {
        const proxyUrl = new URL(notifier.httpProxyUrl);
        if (proxyUrl.protocol !== 'http:' || !proxyUrl.host) {
          return false;
        }
      } catch {
        return false;
      }
    }
  }

  return true;
};
