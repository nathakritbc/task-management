export interface User {
  readonly id: string;
  readonly username: string;
  readonly password?: string | null | undefined;
  readonly full_name: string;
  readonly two_factor_token?: string | null | undefined;
  readonly profile_image?: string | null | undefined;
  readonly invalid_signIn_count?: number | null | undefined;
  readonly status: boolean;
  readonly created_at: Date | string;
  readonly updated_at: Date | string;
}
