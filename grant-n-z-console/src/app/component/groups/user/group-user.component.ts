import {Component, OnInit} from '@angular/core';
import {GroupService} from '../../../service/group.service';
import {ActivatedRoute, Router} from '@angular/router';
import {AppService} from '../../../service/app.service';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {GroupBase} from '../group-base';
import {User} from '../../../model/user';

@Component({
  selector: 'app-group-user',
  templateUrl: './group-user.component.html',
  styleUrls: ['./group-user.component.css']
})
export class GroupUserComponent extends GroupBase implements OnInit {

  public users: User[];
  public displayedColumns: string[];

  /**
   * Constructor.
   *
   * @param appService AppService
   * @param groupService GroupService
   * @param router Router
   * @param activatedRoute ActivatedRoute
   * @param overlay Overlay
   * @param toastrService ToastrService
   */
  constructor(private groupService: GroupService,
              private router: Router,
              public appService: AppService,
              public activatedRoute: ActivatedRoute,
              public overlay: Overlay,
              public toastrService: ToastrService) {
    super(appService, activatedRoute, overlay, toastrService);
  }

  ngOnInit(): void {
    this.showProgress();

    this.groupService.getGroupUserById(this.groupUuid)
      .then(result => {
        this.users = result;
        this.displayedColumns = ['uuid', 'email', 'username'];
        this.hideProgress();
      }).catch(_ => {
      this.showErrorMsg('Could not read data.');
      this.hideProgress();
    });
  }
}
