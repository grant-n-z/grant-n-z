import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {ServiceService} from '../../../service/service.service';
import {Service} from '../../../model/service';
import {ServiceBase} from '../service-base';
import {AppService} from '../../../service/app.service';

@Component({
  selector: 'app-service-new',
  templateUrl: './service-new.component.html',
  styleUrls: ['./service-new.component.css']
})
export class ServiceNewComponent extends ServiceBase implements OnInit {

  public serviceModel: Service = new Service();
  public submitError = '';

  /**
   * Constructor.
   *
   * @param service ServiceService
   * @param router Router
   * @param appService AppService
   * @param overlay Overlay
   * @param toastrService ToastrService
   */
  constructor(private service: ServiceService,
              private router: Router,
              public appService: AppService,
              public overlay: Overlay,
              public toastrService: ToastrService) {
    super(appService, overlay, toastrService);
  }

  ngOnInit(): void {
  }

  async onSubmit() {
    this.submitError = '';
    if (this.serviceModel.name === undefined || this.serviceModel.name === '') {
      this.submitError = 'Service name is empty';
    }

    this.showProgress();
    const result = await this.service.create(this.serviceModel);
    this.hideProgress();
    if (result) {
      this.showSuccessMsg('Created service');
      return;
    }

    this.showErrorMsg('Failed to create service');
  }
}
