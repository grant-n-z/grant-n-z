import {Injectable} from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from '@angular/router';
import {Observable} from 'rxjs';
import {LocalStorageService} from '../service/infra/local-storage.service';

@Injectable()
export class LoginRequireGuard implements CanActivate {

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
      return true;
    } else {
      this.router.navigate(['/']);
      return false;
    }
  }
}
