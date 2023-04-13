import {Injectable} from '@angular/core';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {ComponentPortal} from '@angular/cdk/portal';
import {MatSpinner} from '@angular/material/progress-spinner';
import {AppService} from '../../service/app.service';

@Injectable()
export class ServiceBase {
  public progress = this.overlay.create({
    hasBackdrop: true,
    positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically()
  });

  constructor(public appService: AppService,
              public overlay: Overlay,
              public toastrService: ToastrService) {

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
