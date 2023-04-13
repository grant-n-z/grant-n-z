export class JwtPayload {
  exp: string;
  iat: string;
  iss: string;
  sub: string;
  // tslint:disable-next-line:variable-name
  policy_id: string;
  // tslint:disable-next-line:variable-name
  role_id: string;
  // tslint:disable-next-line:variable-name
  service_id: string;
  username: string;
}
