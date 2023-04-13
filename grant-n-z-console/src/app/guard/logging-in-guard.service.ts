import {Injectable} from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from '@angular/router';
import {Observable} from 'rxjs';
import {LocalStorageService} from '../service/infra/local-storage.service';

@Injectable()
export class LoggingInGuard implements CanActivate {

  /**
   * Constructor.
   *
   * @param localStorageService LocalStorageService
   * @param router Router
   */
  constructor(private localStorageService: LocalStorageService,
              private router: Router) {
  }

  /**
   * Check authentication.
   *
   * @param next ActivatedRouteSnapshot
   * @param state RouterStateSnapshot
   */
  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> | boolean {
    const auth = this.localStorageService.getAuthCookie();
    if (auth != null && auth !== '') {
      this.router.navigate(['/users']);
      return false;
    } else {
      return true;
    }
  }
}
