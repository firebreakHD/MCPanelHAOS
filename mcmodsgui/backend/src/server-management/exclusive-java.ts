import { ConflictException } from '@nestjs/common';

export const EXCLUSIVE_MC_CONTAINER = process.env.MINEPANEL_HAOS === 'true' ? 'mcpanelhaos-mc' : 'minepanel-exclusive-mc';
export const HAOS_MC_LABEL = 'io.mcpanelhaos.managed';
export const EXCLUSIVE_MC_PORT = '25565';
export const EXCLUSIVE_MC_PROFILES = ['modpack-a', 'modpack-b'] as const;

export function exclusiveJavaEnabled(): boolean {
  return process.env.MINEPANEL_EXCLUSIVE_JAVA === 'true';
}

export function assertExclusiveProfile(id: string): void {
  if (!EXCLUSIVE_MC_PROFILES.some((profile) => profile === id)) {
    throw new ConflictException('Only modpack-a and modpack-b are available in exclusive Java mode');
  }
}

export class ExclusiveJavaOperation {
  private pending?: Promise<unknown>;
  private closing = false;

  async drain(): Promise<void> {
    this.closing = true;
    await this.pending?.catch(() => undefined);
  }

  async run<T>(operation: () => Promise<T>): Promise<T> {
    if (this.closing || this.pending) {
      throw new ConflictException({ code: 'EXCLUSIVE_BUSY', message: 'A Minecraft lifecycle action is already in progress' });
    }
    const result = Promise.resolve().then(operation);
    this.pending = result;
    try {
      return await result;
    } finally {
      this.pending = undefined;
    }
  }
}
