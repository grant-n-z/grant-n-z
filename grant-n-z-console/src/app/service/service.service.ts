import {Injectable} from '@angular/core';
import {environment} from '../../environments/environment';
import {Service} from '../model/service';
import {ApiClientService} from './infra/api-client.service';
import {LocalStorageService} from './infra/local-storage.service';

@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  /**
   * Constructor.
   *
   * @param apiClientService ApiClientService
   * @param localStorageService LocalStorageService
   */
  constructor(private apiClientService: ApiClientService,
              private localStorageService: LocalStorageService) {
  }

  public async getAll(): Promise<any> {
    return await this.apiClientService.get(environment.api_base_url + '/api/v1/services', {})
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to getGroupsOfUser all services.', error);
      });
  }

  public async create(service: Service): Promise<boolean> {
    return await this.apiClientService.post(
      environment.api_base_url + '/api/operators/service', service, this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return !(result === undefined || result === null);
      })
      .catch(_ => {
        return false;
      });
  }

  public async getOfUser(): Promise<any> {
    return await this.apiClientService.get(environment.api_base_url + '/api/v1/users/service', this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to getGroupsOfUser services of user.', error);
      });
  }

  public extractApiKey(services: Array<Service>, selectedName: string): string {
    let clientSecret = '';
    services.forEach(service => {
      if (service.name === selectedName) {
        clientSecret = service.secret;
      }
    });
    return clientSecret;
  }

  public getSecret(): string {
    return this.localStorageService.getClientSecretCookie();
  }
}
