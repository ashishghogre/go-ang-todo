import { Component, OnInit } from "@angular/core";
import { ItemService } from "src/app/services/item.service";

@Component({
  selector: "app-side-drawer",
  templateUrl: "./side-drawer.component.html",
  styleUrls: ["./side-drawer.component.scss"]
})
export class SideDrawerComponent implements OnInit {
  constructor(private itemService: ItemService) {}
  ngOnInit() {}

  createItem() {
    console.log(this.itemService.itemAdded);
    if (!this.itemService.itemAdded) {
      this.itemService.createItem();
      this.itemService.itemAdded = true;
    }
  }
}
