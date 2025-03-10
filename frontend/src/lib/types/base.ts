export interface ISiteItem {
  site_title: string;
  site_remark: string;
  site_command: string;
  site_folder: string;
  site_clone: string;
  site_deploy: string;
  web_url: string;
  git_repository: string;
  git_branch: string;
  git_token: string;
  type: string;
}

export interface ISite {
  sites: ISiteItem[];
}