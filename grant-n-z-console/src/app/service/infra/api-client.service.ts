import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {LocalStorageService} from './local-storage.service';
import {Router} from '@angular/router';
import {environment} from '../../../environments/environment';
import {RefreshTokenRequest} from '../../model/refresh-token-request';

@Injectable({
  providedIn: 'root'
})
export class ApiClientService {
  private retryCnt = 3;

  /**
   * Constructor.
   *
   * @param localStorageService LocalStorageService
   * @param http HttpClient
   * @param router Router
   */
  constructor(private localStorageService: LocalStorageService,
              private http: HttpClient,
              private router: Router) {
  }

  public getGetNoAuthHeaders(clientSecret = null) {
    clientSecret = (clientSecret === null) ? this.localStorageService.getClientSecretCookie() : clientSecret;
    return {
      headers: new HttpHeaders({
        'Client-Secret': clientSecret,
      })
    };
  }

  public getGetAuthHeaders(clientSecret = null, token = null) {
    clientSecret = (clientSecret === null) ? this.localStorageService.getClientSecretCookie() : clientSecret;
    token = (token === null) ? this.localStorageService.getAuthCookie() : token;
    return {
      headers: new HttpHeaders({
        'Client-Secret': (clientSecret === null) ? this.localStorageService.getClientSecretCookie() : clientSecret,
        Authorization: 'Bearer ' + token,
      })
    };
  }

  public getPostNoAuthHeaders(clientSecret = null) {
    clientSecret = (clientSecret === null) ? this.localStorageService.getClientSecretCookie() : clientSecret;
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Client-Secret': clientSecret,
      })
    };
  }

  public getPostAuthHeaders(clientSecret = null, token = null) {
    clientSecret = (clientSecret === null) ? this.localStorageService.getClientSecretCookie() : clientSecret;
    token = (token === null) ? this.localStorageService.getAuthCookie() : token;
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Client-Secret': clientSecret,
        Authorization: 'Bearer ' + token,
      })
    };
  }

  public async post(url, body, options): Promise<any> {
    let response: any = null;
    let index = 0;

    while (true) {
      index++;
      response = await this.http.post(url, body, options)
        .pipe()
        .toPromise()
        .then(result => {
          return result;
        })
        .catch(error => {
          console.log('Failed to POST request.', error);
          return this.errorHandling(error, 'post').then(headers => {
            if (headers === null) {
              throw new Error(error.toString());
            }
            options = headers;
            return 'retry';
          });
        });

      if (response === 'retry' && this.retryCnt > index) {
        continue;
      }
      break;
    }
    return response;
  }

  public async get(url, options): Promise<any> {
    let response: any;
    let index = 0;

    while (true) {
      index++;
      response = await this.http.get(url, options)
        .pipe()
        .toPromise()
        .then(result => {
          return result;
        })
        .catch(error => {
          console.log('Failed to GET request.', error);
          return this.errorHandling(error, 'get').then(headers => {
            if (headers === null) {
              throw new Error(error.toString());
            }
            options = headers;
            return 'retry';
          });
        });

      if (response === 'retry' && this.retryCnt > index) {
        continue;
      }
      break;
    }
    return response;
  }

  public delete() {
  }

  public put() {
  }

  private async errorHandling(error, method): Promise<any> {
    if (error.status === 401) {
      if (JSON.parse(JSON.stringify(error.error)).message.includes('expired')) {
        return await this.refreshToken(method).then(headers => {
          return headers;
        }).catch(err => {
          console.log('Error auto refresh.', err);
          // this.localStorageService.clearCookie();
          // this.router.navigate(['/']);
        });
      }

      // this.localStorageService.clearCookie();
      // this.router.navigate(['/']);
    }
    return null;
  }

  private async refreshToken(method): Promise<any> {
    const refreshTokenRequest = new RefreshTokenRequest();
    refreshTokenRequest.grant_type = 'refresh_token';
    refreshTokenRequest.refresh_token = this.localStorageService.getAuthRCookie();
    return await this.http.post(environment.api_base_url + '/api/v1/token', refreshTokenRequest, this.getPostNoAuthHeaders())
      .pipe()
      .toPromise()
      .then(result => {
        console.log('Token refresh.');
        const body = JSON.parse(JSON.stringify(result));
        this.localStorageService.setAuthCookie(body.token);
        this.localStorageService.setAuthRCookie(body.refresh_token);

        if (method === 'get') {
          return this.getGetAuthHeaders(null, body.token);
        } else if (method === 'post') {
          return this.getPostAuthHeaders(null, body.token);
        }
      }).catch(err => {
        throw err;
      });
  }
}
