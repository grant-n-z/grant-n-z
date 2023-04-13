export class LoginRequest {
  email: string;
  password: string;

  /**
   * Validate email expression.
   */
  public emailPattern(): string {
    return '^[A-Za-z0-9+]+[\\w-]@[\\w\\.-]+\\.\\w{2,}$';
  }

  /**
   * Validate min password.
   */
  public passwordMin(): number {
    return 8;
  }
}
