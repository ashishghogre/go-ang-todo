import { Component, OnInit } from '@angular/core';
import { Item } from '../../models/item';
import { ItemService } from '../../services/item.service';

@Component({
  selector: 'app-body',
  templateUrl: './body.component.html',
  styleUrls: ['./body.component.scss']
})
export class BodyComponent implements OnInit {

  constructor(private itemService: ItemService) {}

  ngOnInit() {
    this.itemService.getItems();
  }

  getItems(): Item[]{
    return ItemService.getAllItems();
  }
}
