export interface Tasks {
  readonly id: string;
  readonly title: string;
  readonly description: string;
  readonly status: string;
  readonly created_at: Date | string;
  readonly updated_at: Date | string;
}

export interface CreateTasks {
  readonly id?: string | null;
  readonly title: string;
  readonly description: string;
  readonly status: string;
  readonly created_at?: Date | string | null;
  readonly updated_at?: Date | string | null;
}
