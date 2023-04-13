import {Injectable} from '@angular/core';
import {CookieService} from 'ngx-cookie-service';
import {environment} from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LocalStorageService {
  static AUTH_COOKIE = 'GRANTNZ_A_ID';
  static AUTH_R_COOKIE = 'GRANTNZ_A_ID_R';
  static CLIENT_SECRET = 'CLIENT_SECRET';
  static GROUP_ID = 'GID';

  /**
   * Constructor.
   *
   * @param cookieService CookieService
   */
  constructor(private cookieService: CookieService) {
  }

  public clearCookie() {
    this.cookieService.deleteAll();
  }

  public setAuthCookie(token: string) {
    this.cookieService.set(LocalStorageService.AUTH_COOKIE, token, null, '/', environment.hostname, false, 'Strict');
  }

  public getAuthCookie(): string {
    return this.cookieService.get(LocalStorageService.AUTH_COOKIE);
  }

  public setAuthRCookie(token: string) {
    this.cookieService.set(LocalStorageService.AUTH_R_COOKIE, token, null, '/', environment.hostname, false, 'Strict');
  }

  public getAuthRCookie(): string {
    return this.cookieService.get(LocalStorageService.AUTH_R_COOKIE);
  }

  public setClientSecretCookie(ClientSecret: string) {
    this.cookieService.set(LocalStorageService.CLIENT_SECRET, ClientSecret, null, '/', environment.hostname, false, 'Strict');
  }

  public getClientSecretCookie(): string {
    return this.cookieService.get(LocalStorageService.CLIENT_SECRET);
  }

  public setGroupIdCookie(groupUuid: string) {
    this.cookieService.set(LocalStorageService.GROUP_ID, groupUuid, null, '/', environment.hostname, false, 'Strict');
  }

  public getGroupIdCookie(): string {
    return this.cookieService.get(LocalStorageService.GROUP_ID);
  }

  public getUsername(): string {
    const token = this.cookieService.get(LocalStorageService.AUTH_COOKIE);
    if (token === null || token === '') {
      return null;
    }

    const payload = token.split('.')[1];
    const authUser = atob(payload);
    return JSON.parse(authUser).username;
  }
}
