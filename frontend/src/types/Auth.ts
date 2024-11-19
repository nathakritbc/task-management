export interface AuthSignIn {
  readonly username: string;
  readonly password: string;
  readonly otp_code?: string | null;
}

export interface UserSignIn {
  readonly user_id: string;
  readonly username: string;
  readonly token: string;
}

export interface AuthSignUp {
  readonly username: string;
  readonly password: string;
  readonly full_name: string;
  readonly confirmPassword?: string | null | undefined;
  readonly profile_image?: string | null | undefined;
}
