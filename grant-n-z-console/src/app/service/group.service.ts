import {Injectable} from '@angular/core';
import {ApiClientService} from './infra/api-client.service';
import {environment} from '../../environments/environment';
import {LocalStorageService} from './infra/local-storage.service';

@Injectable({
  providedIn: 'root'
})
export class GroupService {

  /**
   * Constructor.
   *
   * @param apiClientService ApiClientService
   * @param localStorageService LocalStorageService
   */
  constructor(private apiClientService: ApiClientService,
              private localStorageService: LocalStorageService) {
  }

  public async getGroupsOfUser(): Promise<any> {
    return await this.apiClientService.get(
      environment.api_base_url + '/api/v1/users/group', this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to get groups of user.');
        throw new Error(error);
      });
  }

  public async getGroupById(groupId: string): Promise<any> {
    return await this.apiClientService.get(
      `${environment.api_base_url}/api/v1/groups/${groupId}`, this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to get group by id.');
        throw new Error(error);
      });
  }

  public async getGroupUserById(groupId: string): Promise<any> {
    return await this.apiClientService.get(
      `${environment.api_base_url}/api/v1/groups/${groupId}/user`, this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to get group user by id.');
        throw new Error(error);
      });
  }

  public async getGroupPolicyById(groupId: string): Promise<any> {
    return await this.apiClientService.get(
      `${environment.api_base_url}/api/v1/groups/${groupId}/policy`, this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to get group policy by id.');
        throw new Error(error);
      });
  }

  public async getGroupRoleById(groupId: string): Promise<any> {
    return await this.apiClientService.get(
      `${environment.api_base_url}/api/v1/groups/${groupId}/role`, this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to get group role by id.');
        throw new Error(error);
      });
  }

  public async getGroupPermissionById(groupId: string): Promise<any> {
    return await this.apiClientService.get(
      `${environment.api_base_url}/api/v1/groups/${groupId}/permission`, this.apiClientService.getGetAuthHeaders())
      .then(result => {
        return result;
      })
      .catch(error => {
        console.log('Failed to get group permission by id.' + error);
        return error;
      });
  }

  public updateGid(groupUuid: string) {
    this.localStorageService.setGroupIdCookie(groupUuid);
  }
}
