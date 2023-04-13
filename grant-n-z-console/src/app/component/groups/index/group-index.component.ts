import {Component, OnInit} from '@angular/core';
import {GroupBase} from '../group-base';
import {GroupService} from '../../../service/group.service';
import {Overlay} from '@angular/cdk/overlay';
import {ToastrService} from 'ngx-toastr';
import {Group} from '../../../model/group';
import {ActivatedRoute} from '@angular/router';
import {AppService} from '../../../service/app.service';

@Component({
  selector: 'app-index',
  templateUrl: './group-index.component.html',
  styleUrls: ['./group-index.component.css']
})
export class GroupIndexComponent extends GroupBase implements OnInit {
  public group: Group = new Group();

  /**
   * Constructor.
   *
   * @param appService AppService
   * @param groupService GroupService
   * @param overlay Overlay
   * @param toastrService ToastrService
   * @param activatedRoute ActivatedRoute
   */
  constructor(private groupService: GroupService,
              public appService: AppService,
              public activatedRoute: ActivatedRoute,
              public overlay: Overlay,
              public toastrService: ToastrService) {
    super(appService, activatedRoute, overlay, toastrService);
  }

  ngOnInit(): void {
    this.showProgress();

    this.groupService.getGroupById(this.groupUuid)
      .then(result => {
        this.group = result;
        this.hideProgress();
      }).catch(_ => {
      this.showErrorMsg('Could not read data.');
      this.hideProgress();
    });
  }
}
