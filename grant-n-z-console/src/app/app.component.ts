import {Component} from '@angular/core';
import {environment} from '../environments/environment';
import {ActivatedRoute, Router} from '@angular/router';
import {MatDialog} from '@angular/material/dialog';
import {DialogListComponent} from './component/dialog/dialog-list.component';
import {UserService} from './service/user.service';
import {ServiceService} from './service/service.service';
import {ToastrService} from 'ngx-toastr';
import {Overlay} from '@angular/cdk/overlay';
import {ComponentPortal} from '@angular/cdk/portal';
import {MatSpinner} from '@angular/material/progress-spinner';
import {RefreshTokenRequest} from './model/refresh-token-request';
import {GroupService} from './service/group.service';
import {AppService} from './service/app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public static groupId: string;
  public consoleName = environment.name;
  public username: string;
  public navData = new Map<string, string>();
  public progress = this.overlay.create({
    hasBackdrop: true,
    positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically()
  });

  /**
   * Constructor.
   *
   * @param appService AppService
   * @param userService UserService
   * @param service ServiceService
   * @param groupService GroupService
   * @param toastrService ToastrService
   * @param activatedRoute ActivatedRoute
   * @param router Router
   * @param overlay Overlay
   * @param dialog MatDialog
   */
  constructor(private appService: AppService,
              private userService: UserService,
              private service: ServiceService,
              private groupService: GroupService,
              private toastrService: ToastrService,
              private activatedRoute: ActivatedRoute,
              private router: Router,
              private overlay: Overlay,
              public dialog: MatDialog) {

    this.load();
  }

  openNewService(): void {
    this.router.navigate(['/services/new']);
  }

  openServiceDialog(): void {
    this.service.getOfUser()
      .then(result => {
        const dialogRef = this.dialog.open(DialogListComponent, {
          width: '600px',
          data: {
            title: 'Select a service (project)',
            displayedColumns: ['current', 'name', 'selection'],
            clientSecret: this.service.getSecret(),
            data: result,
          },
          panelClass: 'dialog-list'
        });
        dialogRef.afterClosed().subscribe(next => {
          if (next === null || next === undefined) {
            return;
          }
          this.changeSelectedService(next);
        });
      }).catch(_ => {
    });
  }

  onLogout() {
    this.userService.logout();
    this.toastrService.success('Success logout');
    this.router.navigate(['/']);
  }

  private load() {
    this.appService.subscribeNavMenu().subscribe(next => {
      this.navData = next;
    });

    this.username = this.userService.getUserName();
    if (this.username === null) {
      this.router.navigate(['/']);
    }
  }

  private changeSelectedService(clientSecret: string) {
    this.showProgress();
    const request = new RefreshTokenRequest();
    request.grant_type = 'refresh_token';
    request.refresh_token = this.userService.getAuthRCookie();
    this.userService.auth(request, clientSecret)
      .then(_ => {
        this.router.navigate(['/users']);
        this.toastrService.success('Update service');
        this.hideProgress();
      })
      .catch(_ => {
        this.toastrService.error('Failed to update service');
        this.hideProgress();
      });
  }

  private showProgress(): void {
    this.progress.attach(new ComponentPortal(MatSpinner));
  }

  private hideProgress(): void {
    this.progress.detach();
  }
}
