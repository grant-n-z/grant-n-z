import {Injectable} from '@angular/core';
import {environment} from '../../environments/environment';
import {LocalStorageService} from './infra/local-storage.service';
import {ApiClientService} from './infra/api-client.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  /**
   * Constructor.
   *
   * @param localStorageService LocalStorageService
   * @param apiClientService ApiClientService
   */
  constructor(private localStorageService: LocalStorageService,
              private apiClientService: ApiClientService) {
  }

  public async auth(body, clientSecret: string): Promise<any> {
    return this.apiClientService.post(
      environment.api_base_url + '/api/v1/token?type=operator', body, this.apiClientService.getPostNoAuthHeaders(clientSecret))
      .then(result => {
        const response = JSON.parse(JSON.stringify(result));
        this.localStorageService.setAuthCookie(response.token);
        this.localStorageService.setAuthRCookie(response.refresh_token);
        this.localStorageService.setClientSecretCookie(clientSecret);
      });
  }

  public getAuthCookie(): string {
    return this.localStorageService.getAuthCookie();
  }

  public getAuthRCookie(): string {
    return this.localStorageService.getAuthRCookie();
  }

  public getUserName(): string {
    return this.localStorageService.getUsername();
  }

  public logout() {
    this.localStorageService.clearCookie();
  }
}
