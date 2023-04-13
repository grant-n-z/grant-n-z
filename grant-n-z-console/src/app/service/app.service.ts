import { Injectable } from '@angular/core';
import {BehaviorSubject} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  private navMenu: BehaviorSubject<Map<string, string>> = new BehaviorSubject<Map<string, string>>(null);

  /**
   * Constructor.
   */
  constructor() {
    this.navMenu.next(new Map<string, string>().set('Please Login', '/'));
  }

  public updateNavMenu(isUser: boolean, groupId = '0') {
    const menu = new Map<string, string>();
    if (isUser) {
      menu.set('Policy', `/users/policy`);
      menu.set('Group', `/users`);
      this.navMenu.next(menu);
    } else {
      menu.set('Permission', `/groups/${groupId}/permission`);
      menu.set('Role', `/groups/${groupId}/role`);
      menu.set('Policy', `/groups/${groupId}/policy`);
      menu.set('User', `/groups/${groupId}/user`);
      menu.set('Group', `/groups/${groupId}`);
      this.navMenu.next(menu);
    }
  }

  public subscribeNavMenu() {
    return this.navMenu;
  }
}
