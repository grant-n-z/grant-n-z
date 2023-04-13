import {Component, OnInit} from '@angular/core';
import {UserService} from '../../service/user.service';
import {LoginRequest} from '../../model/login-request';
import {Overlay} from '@angular/cdk/overlay';
import {ComponentPortal} from '@angular/cdk/portal';
import {MatSpinner} from '@angular/material/progress-spinner';
import {ToastrService} from 'ngx-toastr';
import {Router} from '@angular/router';
import {environment} from '../../../environments/environment';
import {ServiceService} from '../../service/service.service';
import {Service} from '../../model/service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public services: Array<Service>;
  public selectedServiceName: string;
  public loginRequest: LoginRequest = new LoginRequest();
  public loginError: string;

  public name = environment.name;
  public progress = this.overlay.create({
    hasBackdrop: true,
    positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically()
  });

  /**
   * Constructor.
   *
   * @param loginService UserService
   * @param service ServiceService
   * @param router Router
   * @param overlay Overlay
   * @param toastrService ToastrService
   */
  constructor(private loginService: UserService,
              private service: ServiceService,
              private router: Router,
              private overlay: Overlay,
              private toastrService: ToastrService) {
  }

  ngOnInit(): void {
    this.service.getAll().then(result => {
      this.services = result;
    });
  }

  onSubmit() {
    this.loginError = '';
    this.showProgress();
    const clientSecret = this.service.extractApiKey(this.services, this.selectedServiceName);
    this.loginService.auth(this.loginRequest, clientSecret)
      .finally(() => {
        this.hideProgress();
        const token = this.loginService.getAuthCookie();
        if (token == null || token === '') {
          this.loginError = 'Email or Password is invalid';
        } else {
          this.toastrService.success('Success auth');
          window.location.reload();
          this.router.navigate(['/users']);
        }
      });
  }

  private showProgress(): void {
    this.progress.attach(new ComponentPortal(MatSpinner));
  }

  private hideProgress(): void {
    this.progress.detach();
  }
}
