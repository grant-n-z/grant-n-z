import {Component, OnInit} from '@angular/core';
import {GroupService} from '../../../service/group.service';
import {Group} from '../../../model/group';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {UserBase} from '../user-base';
import {Router} from '@angular/router';
import {AppService} from '../../../service/app.service';

@Component({
  selector: 'app-home',
  templateUrl: './user-index.component.html',
  styleUrls: ['./user-index.component.css']
})
export class UserIndexComponent extends UserBase implements OnInit {
  public groups: Group[];
  public displayedColumns: string[];

  /**
   * Constructor.
   *
   * @param appService AppService
   * @param groupService GroupService
   * @param router Router
   * @param overlay Overlay
   * @param toastrService ToastrService
   */
  constructor(private groupService: GroupService,
              private router: Router,
              public appService: AppService,
              public overlay: Overlay,
              public toastrService: ToastrService) {
    super(appService, overlay, toastrService);
  }

  ngOnInit(): void {
    this.showProgress();

    this.groupService.getGroupsOfUser()
      .then(result => {
        console.log(result);
        this.groups = result;
        this.displayedColumns = ['id', 'name', 'uuid', 'selection'];
        this.hideProgress();
      }).catch(_ => {
      this.showErrorMsg('Could not read data.');
      this.hideProgress();
    });
  }

  onClickGroup(groupUuid: string): void {
    this.groupService.updateGid(groupUuid);
    this.router.navigate([`/groups/${groupUuid}`]);
  }
}
