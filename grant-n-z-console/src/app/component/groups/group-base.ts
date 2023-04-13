import {Injectable} from '@angular/core';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {ComponentPortal} from '@angular/cdk/portal';
import {MatSpinner} from '@angular/material/progress-spinner';
import {AppService} from '../../service/app.service';
import {ActivatedRoute} from '@angular/router';

@Injectable()
export class GroupBase {
  public groupUuid: string;
  public progress = this.overlay.create({
    hasBackdrop: true,
    positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically()
  });

  constructor(public appService: AppService,
              public activatedRoute: ActivatedRoute,
              public overlay: Overlay,
              public toastrService: ToastrService) {

    this.activatedRoute.paramMap.subscribe(param => {
      this.groupUuid = param.get('group_id');
      this.updateNavMenu(this.groupUuid);
    });
  }

  private updateNavMenu(groupId: string) {
    this.appService.updateNavMenu(false, groupId);
  }

  public showProgress(): void {
    this.progress.attach(new ComponentPortal(MatSpinner));
  }

  public hideProgress(): void {
    this.progress.detach();
  }

  public showSuccessMsg(msg: string): void {
    this.toastrService.success(msg);
  }

  public showErrorMsg(msg: string): void {
    this.toastrService.error(msg);
  }
}
