import { Component, OnInit } from '@angular/core';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {PolicyService} from '../../../service/policy.service';
import {UserBase} from '../user-base';
import {Policy} from '../../../model/policy';
import {AppService} from '../../../service/app.service';

@Component({
  selector: 'app-policy',
  templateUrl: './user-policy.component.html',
  styleUrls: ['./user-policy.component.css']
})
export class UserPolicyComponent extends UserBase implements OnInit {
  public policies: Policy[];
  public displayedColumns: string[];

  /**
   * Constructor.
   *
   * @param appService AppService
   * @param policyService PolicyService
   * @param overlay Overlay
   * @param toastrService ToastrService
   */
  constructor(public appService: AppService,
              public policyService: PolicyService,
              public overlay: Overlay,
              public toastrService: ToastrService) {
    super(appService, overlay, toastrService);
  }

  ngOnInit(): void {
    this.showProgress();

    this.policyService.getOfUser()
      .then(result => {
        this.policies = result;
        this.displayedColumns = ['policy_name', 'group_name', 'service_name', 'role_name', 'permission_name'];
        this.hideProgress();
      }).catch(_ => {
      this.showErrorMsg('Could not read data.');
      this.hideProgress();
    });
  }
}
